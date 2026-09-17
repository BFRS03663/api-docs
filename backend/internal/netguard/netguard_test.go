package netguard

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIsPrivate(t *testing.T) {
	private := []string{"127.0.0.1", "10.1.2.3", "172.16.5.5", "192.168.0.1", "169.254.169.254", "100.64.1.1", "0.0.0.0", "::1", "fe80::1", "fd00::1"}
	public := []string{"8.8.8.8", "1.1.1.1", "2606:4700::1111", "52.1.2.3"}
	for _, s := range private {
		if !IsPrivate(net.ParseIP(s)) {
			t.Errorf("%s should be private", s)
		}
	}
	for _, s := range public {
		if IsPrivate(net.ParseIP(s)) {
			t.Errorf("%s should be public", s)
		}
	}
	if !IsPrivate(nil) {
		t.Error("nil ip must be treated as private")
	}
}

func TestClientBlocksLoopbackUnlessAllowed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(200) }))
	defer srv.Close()

	_, err := NewHTTPClient(5*time.Second, false).Get(srv.URL)
	if err == nil || !errors.Is(err, ErrPrivateAddress) {
		t.Fatalf("loopback should be blocked, got %v", err)
	}
	resp, err := NewHTTPClient(5*time.Second, true).Get(srv.URL)
	if err != nil {
		t.Fatalf("allowPrivate client failed: %v", err)
	}
	resp.Body.Close()

	if err := Control(false)(context.Background(), "tcp", "8.8.8.8:443", nil); err != nil {
		t.Errorf("public address rejected: %v", err)
	}
}
