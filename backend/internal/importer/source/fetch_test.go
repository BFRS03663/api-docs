package source

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchSpecDirect(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write([]byte("openapi: 3.0.0\n"))
	}))
	defer srv.Close()

	res, err := New(true).Fetch(context.Background(), srv.URL+"/specs/api.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if string(res.Data) != "openapi: 3.0.0\n" || res.Filename != "api.yaml" {
		t.Fatalf("res = %+v", res)
	}
}

func TestFetchPostmanDocsPage(t *testing.T) {
	var collectionHit string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write([]byte(`<!doctype html><html><head>` +
				`<meta name="environmentUID" content="1-env">` +
				`<meta name="ownerId" content="8407119"><meta name="publishedId" content="SzYW1zB2">` +
				`<meta name="versionTagId" content="latest"></head><body></body></html>`))
		case strings.HasPrefix(r.URL.Path, "/api/collections/8407119/SzYW1zB2"):
			collectionHit = r.URL.RawQuery
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"info":{"_postman_id":"x"},"item":[]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	res, err := New(true).Fetch(context.Background(), srv.URL+"/")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(collectionHit, "environment=1-env") || !strings.Contains(collectionHit, "segregateAuth=true") {
		t.Errorf("collection query = %q", collectionHit)
	}
	if res.Filename != "SzYW1zB2.postman_collection.json" || !strings.HasPrefix(string(res.Data), `{"info"`) {
		t.Errorf("res = %+v", res)
	}
}

func TestFetchBlocksPrivateHostsByDefault(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("openapi: 3.0.0")) }))
	defer srv.Close()
	if _, err := New(false).Fetch(context.Background(), srv.URL); err == nil || !strings.Contains(err.Error(), "private") {
		t.Fatalf("loopback fetch should be blocked, got %v", err)
	}
}

func TestFetchRejectsUnknownHTMLAndBadURLs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><body>hello</body></html>"))
	}))
	defer srv.Close()

	if _, err := New(true).Fetch(context.Background(), srv.URL); !errors.Is(err, ErrUnsupported) {
		t.Errorf("html err = %v", err)
	}
	for _, bad := range []string{"ftp://x/y", "not a url", "/relative"} {
		if _, err := New(true).Fetch(context.Background(), bad); err == nil {
			t.Errorf("%q: expected error", bad)
		}
	}
	if _, err := New(true).Fetch(context.Background(), srv.URL+"/missing"); err == nil {
		t.Error("expected error for non-200")
	}
}
