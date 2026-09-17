package httpapi

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/gin-gonic/gin"

	"github.com/shiprocket/apidocs/internal/config"
)

func uiRouter(t *testing.T, withUI bool) http.Handler {
	t.Helper()
	cols, ops := seeded()
	d := Deps{
		Config:      &config.Config{GinMode: gin.TestMode, CORSOrigins: []string{"http://localhost:5173"}},
		Mongo:       fakePinger{},
		Logger:      slog.New(slog.DiscardHandler),
		Collections: cols,
		Operations:  ops,
	}
	if withUI {
		d.UI = fstest.MapFS{
			"index.html":          {Data: []byte("<!doctype html><title>App</title>")},
			"assets/app-abc12.js": {Data: []byte("console.log(1)")},
			"logo.jpg":            {Data: []byte("jpg")},
		}
	}
	return NewRouter(d)
}

func get(t *testing.T, h http.Handler, p, accept string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, p, nil)
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	h.ServeHTTP(rec, req)
	return rec
}

func TestSPAFallbackAndAssets(t *testing.T) {
	r := uiRouter(t, true)

	for _, p := range []string{"/", "/admin", "/search?q=x", "/docs/petstore", "/some/deep/route"} {
		rec := get(t, r, p, "text/html")
		if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "<title>App</title>") {
			t.Errorf("%s: %d %s", p, rec.Code, truncate(rec.Body.String()))
		}
	}
	rec := get(t, r, "/assets/app-abc12.js", "")
	if rec.Code != http.StatusOK || rec.Body.String() != "console.log(1)" || !strings.Contains(rec.Header().Get("Cache-Control"), "immutable") {
		t.Errorf("asset: %d %q cache=%q", rec.Code, rec.Body.String(), rec.Header().Get("Cache-Control"))
	}
	if rec := get(t, r, "/logo.jpg", ""); rec.Code != http.StatusOK || strings.Contains(rec.Header().Get("Cache-Control"), "immutable") {
		t.Errorf("public file: %d cache=%q", rec.Code, rec.Header().Get("Cache-Control"))
	}

	// API paths and the text surface keep their JSON 404s.
	for _, p := range []string{"/api/v1/nope", "/api/v1/collections/missing", "/docs/missing.md"} {
		if rec := get(t, r, p, "text/html"); rec.Code != http.StatusNotFound || !strings.Contains(rec.Body.String(), `"error"`) {
			t.Errorf("%s: %d %s", p, rec.Code, rec.Body)
		}
	}
	// Markdown negotiation still wins over the app shell on the bare docs path.
	if rec := get(t, r, "/docs/petstore", "text/markdown"); rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/markdown") {
		t.Errorf("docs markdown with UI present: %d %s", rec.Code, rec.Header().Get("Content-Type"))
	}
}

func TestSPAWithoutBundle(t *testing.T) {
	r := uiRouter(t, false)
	for _, p := range []string{"/", "/docs/petstore", "/anything"} {
		if rec := get(t, r, p, "text/html"); rec.Code != http.StatusNotFound {
			t.Errorf("%s without bundle: %d", p, rec.Code)
		}
	}
}
