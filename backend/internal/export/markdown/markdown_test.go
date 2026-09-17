package markdown

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
)

var update = flag.Bool("update", false, "rewrite golden files")

func load(t *testing.T, fixture string) (domain.Collection, []byte) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "..", "testdata", fixture))
	if err != nil {
		t.Fatal(err)
	}
	res, err := importer.Import(context.Background(), data)
	if err != nil {
		t.Fatalf("import %s: %v", fixture, err)
	}
	col := domain.Collection{Slug: strings.TrimSuffix(strings.TrimSuffix(fixture, ".yaml"), ".json"), Name: res.Title, Version: res.Version, Description: res.Description, OperationCount: len(res.Operations)}
	return col, res.SpecJSON
}

func golden(t *testing.T, name, got string) {
	t.Helper()
	path := filepath.Join("..", "..", "..", "testdata", "golden", name)
	if *update {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("missing golden %s (run with -update): %v", name, err)
	}
	if string(want) != got {
		t.Errorf("%s differs from golden; run `go test ./internal/export/... -update` after reviewing the diff", name)
	}
}

func TestCollectionMarkdownPetstore(t *testing.T) {
	col, spec := load(t, "petstore-v3.yaml")
	doc, err := Load(spec)
	if err != nil {
		t.Fatal(err)
	}
	md := Collection(col, doc, "https://docs.example.com")

	for _, want := range []string{
		"# Swagger Petstore - OpenAPI 3.0",
		"OpenAPI document: https://docs.example.com/docs/petstore-v3/openapi.json",
		"## pet",
		"### GET /pet/findByStatus",
		"Operation ID: `findPetsByStatus`",
		"| `status` | query |",
		"### POST /pet",
		"#### Request body (application/json)",
		"- `name` (string, required)",
		"**200**",
		"Example (generated from schema):",
		"```json",
		"\"name\": \"doggie\"",
	} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q", want)
		}
	}
	golden(t, "petstore-v3.md", md)
}

func TestOperationAndIndex(t *testing.T) {
	col, spec := load(t, "sample.postman_collection.json")
	doc, err := Load(spec)
	if err != nil {
		t.Fatal(err)
	}
	md, ok := Operation(col, doc, "get-user", "https://docs.example.com")
	if !ok {
		t.Fatal("operation get-user not found")
	}
	for _, want := range []string{"## GET /users/{id}", "| `id` | path | yes |", "**404**", "Full documentation: https://docs.example.com/docs/sample.postman_collection.md"} {
		if !strings.Contains(md, want) {
			t.Errorf("operation markdown missing %q:\n%s", want, md)
		}
	}
	if _, ok := Operation(col, doc, "nope", ""); ok {
		t.Error("unknown operation should not be found")
	}

	full := Collection(col, doc, "")
	if !strings.Contains(full, "Example:\n\n```json") || !strings.Contains(full, "\"email\": \"ada@example.com\"") {
		t.Error("request example not rendered as json block")
	}
	if !strings.Contains(full, "## Users / Admin") {
		t.Error("nested folder display name not used")
	}

	// Postman requests that share an endpoint become named examples.
	dup := []byte(`{"info":{"_postman_id":"x","name":"Dup"},"item":[
	 {"name":"Create order","request":{"method":"POST","header":[{"key":"Content-Type","value":"application/json"}],"url":"https://a.io/orders","body":{"mode":"raw","raw":"{\"kind\":\"basic\"}"}},"response":[]},
	 {"name":"Create order (advanced)","request":{"method":"POST","header":[{"key":"Content-Type","value":"application/json"}],"url":"https://a.io/orders","body":{"mode":"raw","raw":"{\"kind\":\"advanced\"}"}},"response":[]}]}`)
	res, err := importer.Import(context.Background(), dup)
	if err != nil {
		t.Fatal(err)
	}
	dupDoc, _ := Load(res.SpecJSON)
	md = Collection(domain.Collection{Slug: "dup", Name: "Dup"}, dupDoc, "")
	if !strings.Contains(md, "Example (Create order):") || !strings.Contains(md, "Example (Create order (advanced)):") {
		t.Errorf("named examples not rendered:\n%s", md)
	}

	idx := Index("Acme Docs", "https://docs.example.com", []domain.Collection{col, {Slug: "other", Name: "Other", OperationCount: 2}})
	for _, want := range []string{"# Acme Docs", "- [Acme Users API](https://docs.example.com/docs/sample.postman_collection.md): Sample collection", "(6 endpoints)", "llms-full.txt"} {
		if !strings.Contains(idx, want) {
			t.Errorf("index missing %q:\n%s", want, idx)
		}
	}
}
