package importer

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer/postman"
)

func fixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", name))
	if err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}
	return data
}

func TestDetect(t *testing.T) {
	cases := []struct {
		name    string
		data    []byte
		want    Format
		wantErr error
	}{
		{"openapi3 yaml", fixture(t, "petstore-v3.yaml"), FormatOpenAPI3, nil},
		{"swagger2 json", fixture(t, "petstore-v2.json"), FormatOpenAPI2, nil},
		{"postman schema", []byte(`{"info":{"name":"x","schema":"https://schema.getpostman.com/json/collection/v2.1.0/collection.json"},"item":[]}`), FormatPostman, nil},
		{"postman id", []byte(`{"info":{"_postman_id":"abc"},"item":[]}`), FormatPostman, nil},
		{"empty", []byte("   \n"), FormatUnknown, ErrUnknownFormat},
		{"garbage", []byte("{not json: [\n"), FormatUnknown, ErrUnknownFormat},
		{"plain object", []byte(`{"hello":"world"}`), FormatUnknown, ErrUnknownFormat},
		{"array", []byte(`[1,2]`), FormatUnknown, ErrUnknownFormat},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Detect(tc.data)
			if got != tc.want {
				t.Errorf("format = %q, want %q", got, tc.want)
			}
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("err = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func findOp(ops []domain.Operation, method, path string) *domain.Operation {
	for i := range ops {
		if ops[i].Method == method && ops[i].Path == path {
			return &ops[i]
		}
	}
	return nil
}

func TestImportOpenAPI3(t *testing.T) {
	res, err := Import(context.Background(), fixture(t, "petstore-v3.yaml"))
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if res.Source != domain.SourceOpenAPI3 {
		t.Errorf("source = %s", res.Source)
	}
	if res.Title == "" || res.Version == "" {
		t.Errorf("title/version missing: %q %q", res.Title, res.Version)
	}
	if len(res.Servers) == 0 {
		t.Error("servers not extracted")
	}
	if len(res.Operations) == 0 {
		t.Fatal("no operations")
	}
	op := findOp(res.Operations, "GET", "/pet/findByStatus")
	if op == nil {
		t.Fatal("GET /pet/findByStatus missing")
	}
	if op.OperationID != "findPetsByStatus" {
		t.Errorf("operationId = %q", op.OperationID)
	}
	if len(op.Tags) != 1 || op.Tags[0] != "pet" {
		t.Errorf("tags = %v", op.Tags)
	}
	if op.SearchText == "" {
		t.Error("searchText empty")
	}
	for i, o := range res.Operations {
		if o.Order != i {
			t.Fatalf("order not sequential at %d: %d", i, o.Order)
		}
	}
	if len(res.SpecJSON) == 0 {
		t.Error("specJSON empty")
	}
}

func TestImportSwagger2UpgradesAndResolvesRefs(t *testing.T) {
	res, err := Import(context.Background(), fixture(t, "petstore-v2.json"))
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if res.Source != domain.SourceOpenAPI2 {
		t.Errorf("source = %s", res.Source)
	}
	if res.Doc.OpenAPI[:2] != "3." {
		t.Errorf("not upgraded: openapi=%q", res.Doc.OpenAPI)
	}
	if findOp(res.Operations, "GET", "/pet/{petId}") == nil {
		t.Fatal("GET /pet/{petId} missing after upgrade")
	}
	pet := res.Doc.Components.Schemas["Pet"]
	if pet == nil || pet.Value == nil {
		t.Fatal("components.schemas.Pet missing")
	}
	// The 200 response of getPetById references #/components/schemas/Pet; it
	// must be resolved (Value non-nil), not a dangling $ref.
	resp := res.Doc.Paths.Value("/pet/{petId}").Get.Responses.Value("200")
	if resp == nil || resp.Value == nil {
		t.Fatal("200 response missing")
	}
	media := resp.Value.Content.Get("application/json")
	if media == nil || media.Schema == nil || media.Schema.Value == nil {
		t.Fatal("response schema $ref not resolved")
	}
	if res.Servers == nil || len(res.Servers) == 0 {
		t.Errorf("servers not derived from host/basePath: %v", res.Servers)
	}
}

func TestImportRejectsBadInput(t *testing.T) {
	cases := map[string]struct {
		data []byte
		err  error
	}{
		"garbage":          {[]byte("::: nope"), ErrUnknownFormat},
		"openapi no info":  {[]byte("openapi: 3.0.0\npaths: {}\n"), ErrInvalidSpec},
		"openapi no paths": {[]byte("openapi: 3.0.0\ninfo: {title: t, version: '1'}\npaths: {}\n"), ErrInvalidSpec},
		"postman empty":    {[]byte(`{"info":{"_postman_id":"1","name":"n"},"item":[]}`), ErrInvalidSpec},
		"postman v1":       {[]byte(`{"id":"1","name":"old","requests":[]}`), postman.ErrUnsupportedVersion},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := Import(context.Background(), tc.data)
			if !errors.Is(err, tc.err) {
				t.Fatalf("err = %v, want %v", err, tc.err)
			}
		})
	}
}

func TestImportPostman(t *testing.T) {
	res, err := Import(context.Background(), fixture(t, "sample.postman_collection.json"))
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if res.Source != domain.SourcePostman {
		t.Errorf("source = %s", res.Source)
	}
	if res.Title != "Acme Users API" || len(res.Servers) != 2 {
		t.Errorf("title=%q servers=%v", res.Title, res.Servers)
	}
	if len(res.Operations) != 6 {
		t.Errorf("operations = %d, want 6", len(res.Operations))
	}
	op := findOp(res.Operations, "GET", "/users/{id}")
	if op == nil || op.OperationID != "get-user" || op.Tags[0] != "users" {
		t.Fatalf("GET /users/{id} = %+v", op)
	}
}

func TestFallbackOperationID(t *testing.T) {
	if got := fallbackOperationID("GET", "/pet/{petId}/photos"); got != "get_pet_petid_photos" {
		t.Errorf("got %q", got)
	}
}
