package proxy

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// echoUpstream records what the proxy sent and replies with a JSON summary,
// a custom header, and a cookie the proxy must not relay.
func echoUpstream(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Upstream", "yes")
		w.Header().Set("Set-Cookie", "session=upstream")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"method": r.Method,
			"path":   r.URL.Path,
			"query":  r.URL.RawQuery,
			"host":   r.Host,
			"body":   string(body),
			"headers": map[string]string{
				"Authorization": r.Header.Get("Authorization"),
				"X-Custom":      r.Header.Get("X-Custom"),
				"Cookie":        r.Header.Get("Cookie"),
				"Origin":        r.Header.Get("Origin"),
				"Content-Type":  r.Header.Get("Content-Type"),
			},
		})
	}))
}

func router(cfg Config) http.Handler {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Any("/proxy", New(cfg).Gin())
	return r
}

func call(t *testing.T, h http.Handler, method, target, body string, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	var rd io.Reader
	if body != "" {
		rd = strings.NewReader(body)
	}
	req := httptest.NewRequest(method, "/proxy?scalar_url="+url.QueryEscape(target), rd)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	h.ServeHTTP(rec, req)
	return rec
}

func TestForwardsRequestAndRelaysResponse(t *testing.T) {
	up := echoUpstream(t)
	defer up.Close()
	h := router(Config{AllowPrivate: true})

	rec := call(t, h, http.MethodPost, up.URL+"/v1/orders?dry=1", `{"a":1}`, map[string]string{
		"Authorization": "Bearer upstream-token",
		"X-Custom":      "kept",
		"Cookie":        "portal=secret",
		"Origin":        "http://localhost:5173",
		"Content-Type":  "application/json",
		"Connection":    "keep-alive",
	})
	if rec.Code != http.StatusAccepted {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	var echo struct {
		Method, Path, Query, Host, Body string
		Headers                         map[string]string
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &echo); err != nil {
		t.Fatal(err)
	}
	if echo.Method != "POST" || echo.Path != "/v1/orders" || echo.Query != "dry=1" || echo.Body != `{"a":1}` {
		t.Errorf("upstream saw %+v", echo)
	}
	if echo.Headers["Authorization"] != "Bearer upstream-token" || echo.Headers["X-Custom"] != "kept" || echo.Headers["Content-Type"] != "application/json" {
		t.Errorf("headers not forwarded: %v", echo.Headers)
	}
	if echo.Headers["Cookie"] != "" || echo.Headers["Origin"] != "" {
		t.Errorf("origin-bound headers leaked: %v", echo.Headers)
	}
	if !strings.HasPrefix(up.URL, "http://"+echo.Host) {
		t.Errorf("Host header should be the upstream host, got %q", echo.Host)
	}
	if rec.Header().Get("X-Upstream") != "yes" || rec.Header().Get("Set-Cookie") != "" {
		t.Errorf("response headers wrong: %v", rec.Header())
	}
	if rec.Header().Get("Access-Control-Allow-Origin") != "*" || rec.Header().Get("X-Proxy-Target-Host") == "" {
		t.Errorf("cors/target headers missing: %v", rec.Header())
	}
}

func TestPreflight(t *testing.T) {
	h := router(Config{AllowPrivate: true})
	rec := call(t, h, http.MethodOptions, "https://api.example.com/x", "", map[string]string{"Access-Control-Request-Headers": "authorization, x-custom"})
	if rec.Code != http.StatusNoContent || rec.Header().Get("Access-Control-Allow-Headers") != "authorization, x-custom" {
		t.Fatalf("preflight: %d %v", rec.Code, rec.Header())
	}
}

func TestRejectsBadAndForbiddenTargets(t *testing.T) {
	up := echoUpstream(t)
	defer up.Close()

	public := router(Config{})
	cases := map[string]struct {
		h      http.Handler
		target string
		want   int
	}{
		"missing":           {public, "", http.StatusBadRequest},
		"relative":          {public, "/just/a/path", http.StatusBadRequest},
		"ftp":               {public, "ftp://example.com/x", http.StatusBadRequest},
		"credentials":       {public, "https://user:pw@example.com/x", http.StatusBadRequest},
		"loopback literal":  {public, "http://127.0.0.1:9/x", http.StatusForbidden},
		"metadata":          {public, "http://169.254.169.254/latest", http.StatusForbidden},
		"localhost name":    {public, "http://localhost:8080/x", http.StatusForbidden},
		"loopback server":   {public, up.URL + "/x", http.StatusForbidden},
		"outside allowlist": {router(Config{AllowPrivate: true, AllowedHosts: []string{"api.example.com"}}), up.URL + "/x", http.StatusForbidden},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			rec := call(t, tc.h, http.MethodGet, tc.target, "", nil)
			if rec.Code != tc.want {
				t.Fatalf("status %d, want %d: %s", rec.Code, tc.want, rec.Body)
			}
		})
	}

	// Allowlist matches subdomains but not lookalikes.
	if !hostAllowed("v2.api.example.com", []string{"api.example.com"}) || hostAllowed("api.example.com.evil.io", []string{"api.example.com"}) {
		t.Error("allowlist suffix matching wrong")
	}
}

func TestTruncatesLargeResponses(t *testing.T) {
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(strings.Repeat("x", 4096)))
	}))
	defer up.Close()
	h := router(Config{AllowPrivate: true, MaxResponseBytes: 1024})
	rec := call(t, h, http.MethodGet, up.URL, "", nil)
	if rec.Code != http.StatusOK || rec.Body.Len() != 1024 || rec.Header().Get("X-Proxy-Truncated") != "true" {
		t.Fatalf("truncation: %d len=%d hdr=%q", rec.Code, rec.Body.Len(), rec.Header().Get("X-Proxy-Truncated"))
	}
}

func TestUpstreamTimeoutAndFailure(t *testing.T) {
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(400 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer slow.Close()
	h := router(Config{AllowPrivate: true, Timeout: 100 * time.Millisecond})
	if rec := call(t, h, http.MethodGet, slow.URL, "", nil); rec.Code != http.StatusGatewayTimeout {
		t.Errorf("timeout: %d %s", rec.Code, rec.Body)
	}

	dead := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	deadURL := dead.URL
	dead.Close()
	h = router(Config{AllowPrivate: true})
	if rec := call(t, h, http.MethodGet, deadURL, "", nil); rec.Code != http.StatusBadGateway || strings.Contains(rec.Body.String(), "127.0.0.1") {
		t.Errorf("connection refused: %d %s", rec.Code, rec.Body)
	}
}

func TestRequestBodyLimit(t *testing.T) {
	up := echoUpstream(t)
	defer up.Close()
	h := router(Config{AllowPrivate: true, MaxRequestBytes: 16})
	rec := call(t, h, http.MethodPost, up.URL, strings.Repeat("y", 64), nil)
	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("oversize request: %d %s", rec.Code, rec.Body)
	}
}
