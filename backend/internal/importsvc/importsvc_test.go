package importsvc

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
	"github.com/shiprocket/apidocs/internal/importer/source"
)

type memCols struct{ saved map[string]*domain.Collection }

func (m *memCols) Upsert(_ context.Context, c *domain.Collection) (bool, error) {
	if m.saved == nil {
		m.saved = map[string]*domain.Collection{}
	}
	_, exists := m.saved[c.Slug]
	c.ID = "id-" + c.Slug
	cp := *c
	m.saved[c.Slug] = &cp
	return !exists, nil
}

type memOps struct{ replaced map[string]int }

func (m *memOps) Replace(_ context.Context, id string, ops []domain.Operation) error {
	if m.replaced == nil {
		m.replaced = map[string]int{}
	}
	m.replaced[id] = len(ops)
	return nil
}

func petstore(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "petstore-v3.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestImportCreatesThenUpdates(t *testing.T) {
	cols, ops := &memCols{}, &memOps{}
	svc := New(cols, ops, nil)

	col, created, err := svc.Import(context.Background(), Input{Data: petstore(t), Filename: "petstore.yaml", Slug: "petstore"})
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if !created || col.ID != "id-petstore" || col.Name == "" || col.OperationCount == 0 {
		t.Fatalf("created=%v col=%+v", created, col)
	}
	if ops.replaced["id-petstore"] != col.OperationCount {
		t.Errorf("operations not replaced: %v", ops.replaced)
	}
	if col.Source.Filename != "petstore.yaml" || col.Source.Type != domain.SourceOpenAPI3 || len(col.RawUpload) == 0 {
		t.Errorf("source metadata = %+v", col.Source)
	}

	_, created, err = svc.Import(context.Background(), Input{Data: petstore(t), Slug: "petstore", Name: "Renamed"})
	if err != nil || created {
		t.Fatalf("second import: created=%v err=%v", created, err)
	}
	if cols.saved["petstore"].Name != "Renamed" {
		t.Errorf("explicit name not applied: %q", cols.saved["petstore"].Name)
	}
}

func TestImportFromURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write(petstore(t))
	}))
	defer srv.Close()

	svc := New(&memCols{}, &memOps{}, source.New(true))
	col, _, err := svc.Import(context.Background(), Input{URL: srv.URL + "/openapi.yaml", Slug: "remote"})
	if err != nil {
		t.Fatalf("url import: %v", err)
	}
	if col.Source.Filename != "openapi.yaml" {
		t.Errorf("filename = %q", col.Source.Filename)
	}
	if _, _, err := New(&memCols{}, &memOps{}, nil).Import(context.Background(), Input{URL: srv.URL, Slug: "x"}); err == nil {
		t.Error("url import without a fetcher should fail")
	}
}

func TestImportErrors(t *testing.T) {
	svc := New(&memCols{}, &memOps{}, source.New(true))
	cases := map[string]struct {
		in   Input
		want error
	}{
		"bad slug":    {Input{Data: petstore(t), Slug: "Bad Slug"}, ErrInvalidSlug},
		"empty":       {Input{Data: nil, Slug: "ok"}, importer.ErrUnknownFormat},
		"invalid":     {Input{Data: []byte("openapi: 3.0.0\npaths: {}\n"), Slug: "ok"}, importer.ErrInvalidSpec},
		"bad url":     {Input{URL: "not a url", Slug: "ok"}, importer.ErrInvalidSpec},
		"unreachable": {Input{URL: "http://127.0.0.1:1/x", Slug: "ok"}, importer.ErrInvalidSpec},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if _, _, err := svc.Import(context.Background(), tc.in); !errors.Is(err, tc.want) {
				t.Fatalf("err = %v, want %v", err, tc.want)
			}
		})
	}
}
