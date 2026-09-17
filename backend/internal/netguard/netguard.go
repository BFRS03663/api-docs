// Package netguard blocks outbound connections to private, loopback and
// link-local addresses so server-side fetches (URL imports, the Try-it proxy)
// cannot be pointed at internal infrastructure.
package netguard

import (
	"context"
	"errors"
	"net"
	"net/http"
	"syscall"
	"time"
)

// ErrPrivateAddress is returned when a dial targets a non-public address.
var ErrPrivateAddress = errors.New("destination resolves to a private or local address")

// IsPrivate reports whether ip must not be reached from the server.
func IsPrivate(ip net.IP) bool {
	if ip == nil {
		return true
	}
	if ip.IsLoopback() || ip.IsUnspecified() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsMulticast() || ip.IsPrivate() {
		return true
	}
	if v4 := ip.To4(); v4 != nil {
		// Carrier-grade NAT 100.64.0.0/10 and the IPv4-mapped special ranges.
		if v4[0] == 100 && v4[1]&0xc0 == 64 {
			return true
		}
		if v4[0] == 0 || v4[0] == 192 && v4[1] == 0 && v4[2] == 0 {
			return true
		}
	}
	return false
}

// Control returns a dialer control hook that rejects private targets. The
// check runs on the resolved IP, so DNS tricks and redirects are covered.
func Control(allowPrivate bool) func(ctx context.Context, network, address string, c syscall.RawConn) error {
	return func(_ context.Context, _ string, address string, _ syscall.RawConn) error {
		if allowPrivate {
			return nil
		}
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return err
		}
		if IsPrivate(net.ParseIP(host)) {
			return ErrPrivateAddress
		}
		return nil
	}
}

// NewHTTPClient returns a client whose every connection, including those
// made while following redirects, passes through the private-address guard.
func NewHTTPClient(timeout time.Duration, allowPrivate bool) *http.Client {
	dialer := &net.Dialer{Timeout: 10 * time.Second, ControlContext: Control(allowPrivate)}
	transport := &http.Transport{
		DialContext:           dialer.DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
		MaxIdleConns:          10,
		Proxy:                 nil, // never honour HTTP_PROXY for guarded fetches
	}
	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return errors.New("too many redirects")
			}
			return nil
		},
	}
}
