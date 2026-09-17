package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxyRouteIsRateLimited(t *testing.T) {
	r := newTestRouterWith(fakePinger{}, &fakeCols{}, &fakeOps{})
	var last int
	for i := 1; i <= 61; i++ {
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/proxy", nil))
		last = rec.Code
		if i <= 60 && rec.Code != http.StatusBadRequest {
			t.Fatalf("request %d: status %d, want 400 (missing scalar_url)", i, rec.Code)
		}
	}
	if last != http.StatusTooManyRequests {
		t.Fatalf("61st request: status %d, want 429", last)
	}
}

func TestProxyRouteAcceptsAnyMethodAndPreflight(t *testing.T) {
	r := newTestRouterWith(fakePinger{}, &fakeCols{}, &fakeOps{})
	for _, m := range []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch} {
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest(m, "/api/v1/proxy", nil))
		if rec.Code != http.StatusBadRequest {
			t.Errorf("%s: %d, want 400", m, rec.Code)
		}
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/proxy?scalar_url=https%3A%2F%2Fapi.example.com", nil)
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Errorf("preflight: %d", rec.Code)
	}
}
