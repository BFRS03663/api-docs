package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
)

// seededPetstore returns fakes holding the real Petstore spec so Markdown
// rendering runs against a genuine document.
func seededPetstore(t *testing.T) (*fakeCols, *fakeOps, *fakeSearch) {
	t.Helper()
	res, err := importer.Import(context.Background(), petstoreYAML(t))
	if err != nil {
		t.Fatal(err)
	}
	cols := &fakeCols{
		items: map[string]*domain.Collection{"petstore": {ID: "c1", Slug: "petstore", Name: "Petstore", Description: "Pets as a service.\nSecond line.", OperationCount: len(res.Operations)}},
		specs: map[string][]byte{"petstore": res.SpecJSON},
	}
	ops := &fakeOps{byCol: map[string][]domain.Operation{"c1": res.Operations}}
	search := &fakeSearch{ops: res.Operations}
	return cols, ops, search
}

func TestLLMSIndexAndFull(t *testing.T) {
	cols, ops, search := seededPetstore(t)
	r := newRouterWithSearch(cols, ops, search)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/llms.txt", nil)
	req.Host = "docs.example.com"
	req.Header.Set("X-Forwarded-Proto", "https")
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/plain") {
		t.Fatalf("llms.txt: %d %s", rec.Code, rec.Header().Get("Content-Type"))
	}
	body := rec.Body.String()
	for _, want := range []string{"# API Docs", "[Petstore](https://docs.example.com/docs/petstore.md): Pets as a service.", "https://docs.example.com/llms-full.txt"} {
		if !strings.Contains(body, want) {
			t.Errorf("llms.txt missing %q:\n%s", want, body)
		}
	}

	rec = do(t, r, "/llms-full.txt")
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "### GET /pet/findByStatus") {
		t.Errorf("llms-full.txt: %d, contains op: %v", rec.Code, strings.Contains(rec.Body.String(), "findByStatus"))
	}
}

func TestCollectionTextFormats(t *testing.T) {
	cols, ops, search := seededPetstore(t)
	r := newRouterWithSearch(cols, ops, search)

	rec := do(t, r, "/docs/petstore.md")
	if rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/markdown") || !strings.HasPrefix(rec.Body.String(), "# Petstore") {
		t.Errorf(".md: %d %s %s", rec.Code, rec.Header().Get("Content-Type"), truncate(rec.Body.String()))
	}
	rec = do(t, r, "/docs/petstore.txt")
	if rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/plain") {
		t.Errorf(".txt: %d %s", rec.Code, rec.Header().Get("Content-Type"))
	}
	rec = do(t, r, "/docs/petstore/findPetsByStatus.md")
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "## GET /pet/findByStatus") {
		t.Errorf("operation .md: %d %s", rec.Code, truncate(rec.Body.String()))
	}
	if rec := do(t, r, "/docs/petstore/openapi.json"); rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "application/json") {
		t.Errorf("openapi.json: %d", rec.Code)
	}

	// Accept negotiation on the bare path and on the JSON API.
	for _, path := range []string{"/docs/petstore", "/api/v1/collections/petstore"} {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Accept", "text/markdown")
		r.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/markdown") {
			t.Errorf("%s with Accept text/markdown: %d %s", path, rec.Code, rec.Header().Get("Content-Type"))
		}
	}

	for _, path := range []string{"/docs/nope.md", "/docs/petstore/nope.md", "/docs/petstore/photo.png", "/docs/Bad_Slug.md", "/docs/petstore"} {
		if rec := do(t, r, path); rec.Code != http.StatusNotFound {
			t.Errorf("%s: %d, want 404", path, rec.Code)
		}
	}
}

func TestSearchEndpoint(t *testing.T) {
	cols, ops, search := seededPetstore(t)
	r := newRouterWithSearch(cols, ops, search)

	rec := do(t, r, "/api/v1/search?q=status")
	if rec.Code != http.StatusOK {
		t.Fatalf("search: %d %s", rec.Code, rec.Body)
	}
	var body struct {
		Hits []domain.SearchHit `json:"hits"`
	}
	_ = json.Unmarshal(rec.Body.Bytes(), &body)
	if len(body.Hits) == 0 || body.Hits[0].OperationID != "findPetsByStatus" || body.Hits[0].CollectionSlug != "petstore" || body.Hits[0].CollectionName != "Petstore" {
		t.Fatalf("hits = %+v", body.Hits)
	}
	if search.lastCollection != "" || search.lastLimit != 25 {
		t.Errorf("search args = %q %d", search.lastCollection, search.lastLimit)
	}

	if rec := do(t, r, "/api/v1/search?q=status&collection=petstore&limit=5"); rec.Code != http.StatusOK || search.lastCollection != "c1" || search.lastLimit != 5 {
		t.Errorf("filtered search: %d args=%q %d", rec.Code, search.lastCollection, search.lastLimit)
	}
	for _, path := range []string{"/api/v1/search", "/api/v1/search?q=", "/api/v1/search?q=x&limit=0", "/api/v1/search?q=x&limit=abc", "/api/v1/search?q=" + strings.Repeat("a", 201)} {
		if rec := do(t, r, path); rec.Code != http.StatusBadRequest {
			t.Errorf("%s: %d, want 400", path, rec.Code)
		}
	}
	if rec := do(t, r, "/api/v1/search?q=x&collection=missing"); rec.Code != http.StatusNotFound {
		t.Errorf("unknown collection filter: %d", rec.Code)
	}
}
