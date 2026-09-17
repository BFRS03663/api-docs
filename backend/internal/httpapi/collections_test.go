package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shiprocket/apidocs/internal/domain"
)

type fakeCols struct {
	items map[string]*domain.Collection
	specs map[string][]byte
	err   error
}

func (f *fakeCols) List(context.Context) ([]domain.Collection, error) {
	if f.err != nil {
		return nil, f.err
	}
	out := []domain.Collection{}
	for _, c := range f.items {
		out = append(out, *c)
	}
	return out, nil
}

func (f *fakeCols) GetBySlug(_ context.Context, slug string) (*domain.Collection, error) {
	if f.err != nil {
		return nil, f.err
	}
	if c, ok := f.items[slug]; ok {
		return c, nil
	}
	return nil, domain.ErrNotFound
}

func (f *fakeCols) GetSpec(_ context.Context, slug string) ([]byte, error) {
	if s, ok := f.specs[slug]; ok {
		return s, nil
	}
	return nil, domain.ErrNotFound
}

func (f *fakeCols) Upsert(_ context.Context, c *domain.Collection) (bool, error) {
	if f.err != nil {
		return false, f.err
	}
	if f.items == nil {
		f.items = map[string]*domain.Collection{}
	}
	if f.specs == nil {
		f.specs = map[string][]byte{}
	}
	existing, ok := f.items[c.Slug]
	if ok {
		c.ID = existing.ID
	} else {
		c.ID = "id-" + c.Slug
	}
	meta := *c
	meta.Spec, meta.RawUpload = nil, nil
	f.items[c.Slug] = &meta
	f.specs[c.Slug] = c.Spec
	return !ok, nil
}

func (f *fakeCols) Delete(_ context.Context, slug string) (string, error) {
	c, ok := f.items[slug]
	if !ok {
		return "", domain.ErrNotFound
	}
	delete(f.items, slug)
	delete(f.specs, slug)
	return c.ID, nil
}

type fakeOps struct{ byCol map[string][]domain.Operation }

func (f *fakeOps) ListByCollection(_ context.Context, id string) ([]domain.Operation, error) {
	return f.byCol[id], nil
}

func (f *fakeOps) Replace(_ context.Context, id string, ops []domain.Operation) error {
	if f.byCol == nil {
		f.byCol = map[string][]domain.Operation{}
	}
	f.byCol[id] = ops
	return nil
}

func (f *fakeOps) DeleteByCollection(_ context.Context, id string) error {
	delete(f.byCol, id)
	return nil
}

func seeded() (*fakeCols, *fakeOps) {
	cols := &fakeCols{
		items: map[string]*domain.Collection{"petstore": {ID: "c1", Slug: "petstore", Name: "Petstore"}},
		specs: map[string][]byte{"petstore": []byte(`{"openapi":"3.0.0"}`)},
	}
	ops := &fakeOps{byCol: map[string][]domain.Operation{"c1": {{Method: "GET", Path: "/pet", OperationID: "listPets"}}}}
	return cols, ops
}

func do(t *testing.T, h http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestCollectionsList(t *testing.T) {
	cols, ops := seeded()
	rec := do(t, newTestRouterWith(fakePinger{}, cols, ops), "/api/v1/collections")
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d: %s", rec.Code, rec.Body)
	}
	var body struct {
		Collections []domain.Collection `json:"collections"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Collections) != 1 || body.Collections[0].Slug != "petstore" {
		t.Fatalf("body = %s", rec.Body)
	}
}

func TestCollectionGetAndOperations(t *testing.T) {
	cols, ops := seeded()
	r := newTestRouterWith(fakePinger{}, cols, ops)

	if rec := do(t, r, "/api/v1/collections/petstore"); rec.Code != http.StatusOK {
		t.Fatalf("get: %d %s", rec.Code, rec.Body)
	}
	rec := do(t, r, "/api/v1/collections/petstore/operations")
	if rec.Code != http.StatusOK {
		t.Fatalf("ops: %d %s", rec.Code, rec.Body)
	}
	var body struct {
		Operations []domain.Operation `json:"operations"`
	}
	_ = json.Unmarshal(rec.Body.Bytes(), &body)
	if len(body.Operations) != 1 || body.Operations[0].OperationID != "listPets" {
		t.Fatalf("ops body = %s", rec.Body)
	}
}

func TestCollectionNotFoundAndBadSlug(t *testing.T) {
	cols, ops := seeded()
	r := newTestRouterWith(fakePinger{}, cols, ops)
	for _, p := range []string{"/api/v1/collections/nope", "/api/v1/collections/Bad_Slug", "/api/v1/collections/nope/operations", "/docs/nope/openapi.json"} {
		if rec := do(t, r, p); rec.Code != http.StatusNotFound {
			t.Errorf("%s: status %d, want 404", p, rec.Code)
		}
	}
}

func TestSpecServedAsJSON(t *testing.T) {
	cols, ops := seeded()
	rec := do(t, newTestRouterWith(fakePinger{}, cols, ops), "/docs/petstore/openapi.json")
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Errorf("content-type = %q", ct)
	}
	if rec.Body.String() != `{"openapi":"3.0.0"}` {
		t.Errorf("body = %s", rec.Body)
	}
}

func TestStoreErrorIsGeneric(t *testing.T) {
	cols, ops := seeded()
	cols.err = errors.New("connection reset by peer at 10.0.0.5")
	rec := do(t, newTestRouterWith(fakePinger{}, cols, ops), "/api/v1/collections")
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status %d", rec.Code)
	}
	if rec.Body.String() != `{"error":"internal error"}` {
		t.Fatalf("body leaks details: %s", rec.Body)
	}
}
