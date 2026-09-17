// Package proxy forwards "Try it" requests from the docs UI to the
// documented APIs. It implements the contract Scalar's request runner uses:
// the browser calls the proxy with the real destination in the scalar_url
// query parameter and the proxy relays status, headers and body back.
//
// Because the browser gets to pick the destination, the proxy is a
// server-side request forgery surface. Every request passes the private
// address guard (including redirect hops), can be limited to an allowlist
// of hosts, is bounded in time and size, and never carries our own cookies.
package proxy

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/shiprocket/apidocs/internal/netguard"
)

// Config tunes the forwarder.
type Config struct {
	// AllowedHosts restricts destinations to these hosts and their subdomains.
	// Empty means any public host.
	AllowedHosts []string
	// AllowPrivate permits loopback and private destinations (development only).
	AllowPrivate bool
	// Timeout bounds the whole upstream exchange. Zero means 15 seconds.
	Timeout time.Duration
	// MaxResponseBytes caps the relayed body. Zero means 5 MB.
	MaxResponseBytes int64
	// MaxRequestBytes caps the forwarded body. Zero means 10 MB.
	MaxRequestBytes int64
}

// Handler is the Gin handler for the proxy endpoint.
type Handler struct {
	cfg    Config
	client *http.Client
}

// New builds a Handler whose outbound client honours the private-address guard.
func New(cfg Config) *Handler {
	if cfg.Timeout <= 0 {
		cfg.Timeout = 15 * time.Second
	}
	if cfg.MaxResponseBytes <= 0 {
		cfg.MaxResponseBytes = 5 << 20
	}
	if cfg.MaxRequestBytes <= 0 {
		cfg.MaxRequestBytes = 10 << 20
	}
	for i, h := range cfg.AllowedHosts {
		cfg.AllowedHosts[i] = strings.ToLower(strings.TrimSpace(h))
	}
	return &Handler{cfg: cfg, client: netguard.NewHTTPClient(cfg.Timeout, cfg.AllowPrivate)}
}

// hopByHop headers are connection-specific and must not be relayed.
var hopByHop = map[string]bool{
	"Connection": true, "Keep-Alive": true, "Proxy-Authenticate": true, "Proxy-Authorization": true,
	"Proxy-Connection": true, "Te": true, "Trailer": true, "Transfer-Encoding": true, "Upgrade": true,
}

// requestDrop lists incoming headers that belong to our origin, not the target.
var requestDrop = map[string]bool{
	"Host": true, "Cookie": true, "Origin": true, "Referer": true, "Content-Length": true,
	// Let the Go transport negotiate compression so relayed bodies are plain.
	"Accept-Encoding": true,
}

// responseDrop lists upstream headers we never relay to the browser.
var responseDrop = map[string]bool{
	"Set-Cookie": true, "Set-Cookie2": true, "Content-Length": true, "Content-Encoding": true,
	"Access-Control-Allow-Origin": true, "Access-Control-Allow-Credentials": true,
	"Access-Control-Expose-Headers": true, "Access-Control-Allow-Headers": true, "Access-Control-Allow-Methods": true,
}

// Gin returns the handler function. Mount it with router.Any so every method
// and the CORS preflight reach it.
func (h *Handler) Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		setCORS(c)
		if c.Request.Method == http.MethodOptions {
			if req := c.GetHeader("Access-Control-Request-Headers"); req != "" {
				c.Header("Access-Control-Allow-Headers", req)
			}
			c.Status(http.StatusNoContent)
			return
		}
		target, status, msg := h.target(c.Query("scalar_url"))
		if status != 0 {
			c.JSON(status, gin.H{"error": msg})
			return
		}
		h.forward(c, target)
	}
}

// target validates the destination and applies the host allowlist.
func (h *Handler) target(raw string) (*url.URL, int, string) {
	if strings.TrimSpace(raw) == "" {
		return nil, http.StatusBadRequest, "scalar_url query parameter is required"
	}
	u, err := url.Parse(raw)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" || u.User != nil {
		return nil, http.StatusBadRequest, "scalar_url must be an absolute http(s) URL without credentials"
	}
	host := strings.ToLower(u.Hostname())
	if len(h.cfg.AllowedHosts) > 0 && !hostAllowed(host, h.cfg.AllowedHosts) {
		return nil, http.StatusForbidden, "destination host is not in the allowlist"
	}
	if !h.cfg.AllowPrivate {
		if ip := net.ParseIP(host); ip != nil && netguard.IsPrivate(ip) {
			return nil, http.StatusForbidden, "destination is a private or local address"
		}
		if host == "localhost" || strings.HasSuffix(host, ".localhost") || strings.HasSuffix(host, ".local") || strings.HasSuffix(host, ".internal") {
			return nil, http.StatusForbidden, "destination is a private or local address"
		}
	}
	return u, 0, ""
}

func hostAllowed(host string, allowed []string) bool {
	for _, a := range allowed {
		if a == "" {
			continue
		}
		if host == a || strings.HasSuffix(host, "."+a) {
			return true
		}
	}
	return false
}

func (h *Handler) forward(c *gin.Context, target *url.URL) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.Timeout)
	defer cancel()

	var body io.Reader
	if c.Request.Body != nil && c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
		body = http.MaxBytesReader(c.Writer, c.Request.Body, h.cfg.MaxRequestBytes)
	}
	req, err := http.NewRequestWithContext(ctx, c.Request.Method, target.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not build upstream request"})
		return
	}
	for name, values := range c.Request.Header {
		if hopByHop[name] || requestDrop[name] {
			continue
		}
		for _, v := range values {
			req.Header.Add(name, v)
		}
	}
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "apidocs-try-it/1.0")
	}

	resp, err := h.client.Do(req)
	if err != nil {
		var mbe *http.MaxBytesError
		switch {
		case errors.Is(err, netguard.ErrPrivateAddress):
			c.JSON(http.StatusForbidden, gin.H{"error": "destination is a private or local address"})
		case errors.As(err, &mbe):
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "request body exceeds the proxy limit"})
		case errors.Is(err, context.DeadlineExceeded) || isTimeout(err):
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "upstream did not respond in time"})
		default:
			c.JSON(http.StatusBadGateway, gin.H{"error": "upstream request failed"})
		}
		return
	}
	defer resp.Body.Close()

	limited, err := io.ReadAll(io.LimitReader(resp.Body, h.cfg.MaxResponseBytes+1))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || isTimeout(err) {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "upstream did not respond in time"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "upstream response could not be read"})
		return
	}
	truncated := int64(len(limited)) > h.cfg.MaxResponseBytes
	if truncated {
		limited = limited[:h.cfg.MaxResponseBytes]
	}

	for name, values := range resp.Header {
		if hopByHop[name] || responseDrop[name] {
			continue
		}
		for _, v := range values {
			c.Writer.Header().Add(name, v)
		}
	}
	setCORS(c)
	c.Header("X-Proxy-Target-Host", target.Host)
	if truncated {
		c.Header("X-Proxy-Truncated", "true")
	}
	c.Header("Content-Length", strconv.Itoa(len(limited)))
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), bytes.TrimSuffix(limited, nil))
}

func setCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Max-Age", "600")
}

func isTimeout(err error) bool {
	var ne net.Error
	return errors.As(err, &ne) && ne.Timeout()
}
