package httpapi

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func petstoreYAML(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "testdata", "petstore-v3.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func loginToken(t *testing.T, r http.Handler) string {
	t.Helper()
	rec := httptest.NewRecorder()
	body := strings.NewReader(`{"username":"admin","password":"` + testPassword + `"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", body)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("login status %d: %s", rec.Code, rec.Body)
	}
	var out struct {
		Token string `json:"token"`
	}
	_ = json.Unmarshal(rec.Body.Bytes(), &out)
	if out.Token == "" {
		t.Fatal("no token in login response")
	}
	return out.Token
}

func multipartBody(t *testing.T, fields map[string]string, filename string, data []byte) (*bytes.Buffer, string) {
	t.Helper()
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for k, v := range fields {
		_ = w.WriteField(k, v)
	}
	if filename != "" {
		part, err := w.CreateFormFile("file", filename)
		if err != nil {
			t.Fatal(err)
		}
		_, _ = part.Write(data)
	}
	_ = w.Close()
	return &buf, w.FormDataContentType()
}

func send(r http.Handler, method, path, token, contentType string, body io.Reader) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, body)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	r.ServeHTTP(rec, req)
	return rec
}

func TestLoginFailuresAreGenericAndRateLimited(t *testing.T) {
	r := newTestRouterWith(fakePinger{}, &fakeCols{}, &fakeOps{})
	for i := 1; i <= 6; i++ {
		rec := send(r, http.MethodPost, "/api/v1/auth/login", "", "application/json", strings.NewReader(`{"username":"admin","password":"nope"}`))
		switch {
		case i <= 5 && rec.Code != http.StatusUnauthorized:
			t.Fatalf("attempt %d: status %d, want 401", i, rec.Code)
		case i <= 5 && rec.Body.String() != `{"error":"invalid credentials"}`:
			t.Fatalf("attempt %d: body leaks detail: %s", i, rec.Body)
		case i == 6 && rec.Code != http.StatusTooManyRequests:
			t.Fatalf("attempt 6: status %d, want 429", rec.Code)
		}
	}
	rec := send(r, http.MethodPost, "/api/v1/auth/login", "", "application/json", strings.NewReader(`{"username":"admin"}`))
	if rec.Code != http.StatusTooManyRequests && rec.Code != http.StatusBadRequest {
		t.Fatalf("missing password: status %d", rec.Code)
	}
}

func TestAdminRoutesRequireToken(t *testing.T) {
	r := newTestRouterWith(fakePinger{}, &fakeCols{}, &fakeOps{})
	for _, tc := range []struct{ method, path string }{
		{http.MethodPost, "/api/v1/admin/collections"},
		{http.MethodPut, "/api/v1/admin/collections/petstore"},
		{http.MethodDelete, "/api/v1/admin/collections/petstore"},
		{http.MethodGet, "/api/v1/auth/me"},
	} {
		if rec := send(r, tc.method, tc.path, "", "", nil); rec.Code != http.StatusUnauthorized {
			t.Errorf("%s %s without token: %d", tc.method, tc.path, rec.Code)
		}
		if rec := send(r, tc.method, tc.path, "bogus", "", nil); rec.Code != http.StatusUnauthorized {
			t.Errorf("%s %s with bad token: %d", tc.method, tc.path, rec.Code)
		}
	}
}

func TestAdminUploadLifecycle(t *testing.T) {
	cols, ops := &fakeCols{}, &fakeOps{}
	r := newTestRouterWith(fakePinger{}, cols, ops)
	token := loginToken(t, r)

	if rec := send(r, http.MethodGet, "/api/v1/auth/me", token, "", nil); rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "admin") {
		t.Fatalf("me: %d %s", rec.Code, rec.Body)
	}

	body, ct := multipartBody(t, map[string]string{"slug": "petstore", "name": "Petstore"}, "petstore.yaml", petstoreYAML(t))
	rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, ct, body)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create: %d %s", rec.Code, rec.Body)
	}
	if !strings.Contains(rec.Body.String(), `"created":true`) || !strings.Contains(rec.Body.String(), `"slug":"petstore"`) {
		t.Errorf("create body = %s", rec.Body)
	}
	if len(ops.byCol["id-petstore"]) == 0 {
		t.Error("operations not indexed after upload")
	}

	// Public read API sees the upload.
	if rec := send(r, http.MethodGet, "/docs/petstore/openapi.json", "", "", nil); rec.Code != http.StatusOK {
		t.Errorf("spec after upload: %d", rec.Code)
	}

	body, ct = multipartBody(t, map[string]string{"slug": "petstore"}, "petstore.yaml", petstoreYAML(t))
	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, ct, body); rec.Code != http.StatusConflict {
		t.Errorf("duplicate POST: %d %s", rec.Code, rec.Body)
	}

	body, ct = multipartBody(t, map[string]string{"name": "Petstore v2"}, "petstore.yaml", petstoreYAML(t))
	if rec := send(r, http.MethodPut, "/api/v1/admin/collections/petstore", token, ct, body); rec.Code != http.StatusOK {
		t.Errorf("reimport: %d %s", rec.Code, rec.Body)
	}
	if cols.items["petstore"].Name != "Petstore v2" {
		t.Errorf("reimport did not update name: %q", cols.items["petstore"].Name)
	}

	body, ct = multipartBody(t, nil, "petstore.yaml", petstoreYAML(t))
	if rec := send(r, http.MethodPut, "/api/v1/admin/collections/missing", token, ct, body); rec.Code != http.StatusNotFound {
		t.Errorf("reimport unknown: %d", rec.Code)
	}

	if rec := send(r, http.MethodDelete, "/api/v1/admin/collections/petstore", token, "", nil); rec.Code != http.StatusNoContent {
		t.Errorf("delete: %d %s", rec.Code, rec.Body)
	}
	if rec := send(r, http.MethodDelete, "/api/v1/admin/collections/petstore", token, "", nil); rec.Code != http.StatusNotFound {
		t.Errorf("second delete: %d", rec.Code)
	}
	if _, still := ops.byCol["id-petstore"]; still {
		t.Error("operations not cascaded on delete")
	}
}

func TestAdminUploadValidation(t *testing.T) {
	r := newTestRouterWith(fakePinger{}, &fakeCols{}, &fakeOps{})
	token := loginToken(t, r)

	body, ct := multipartBody(t, map[string]string{"slug": "bad"}, "x.yaml", []byte("openapi: 3.0.0\npaths: {}\n"))
	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, ct, body); rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), "invalid specification") {
		t.Errorf("invalid spec: %d %s", rec.Code, rec.Body)
	}

	body, ct = multipartBody(t, map[string]string{"slug": "Bad Slug"}, "x.yaml", petstoreYAML(t))
	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, ct, body); rec.Code != http.StatusBadRequest {
		t.Errorf("bad slug: %d", rec.Code)
	}

	body, ct = multipartBody(t, map[string]string{"slug": "nofile"}, "", nil)
	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, ct, body); rec.Code != http.StatusBadRequest {
		t.Errorf("missing file: %d", rec.Code)
	}

	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, "text/plain", strings.NewReader("x")); rec.Code != http.StatusUnsupportedMediaType {
		t.Errorf("wrong content type: %d", rec.Code)
	}

	huge := bytes.Repeat([]byte("a"), maxUploadBytes+1024)
	body, ct = multipartBody(t, map[string]string{"slug": "huge"}, "huge.yaml", huge)
	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, ct, body); rec.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("oversize: %d %s", rec.Code, truncate(rec.Body.String()))
	}
}

func TestAdminImportFromURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write(petstoreYAML(t))
	}))
	defer srv.Close()

	r := newTestRouterWith(fakePinger{}, &fakeCols{}, &fakeOps{})
	token := loginToken(t, r)
	rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, "application/json", strings.NewReader(`{"url":"`+srv.URL+`/spec.yaml","slug":"remote"}`))
	if rec.Code != http.StatusCreated {
		t.Fatalf("url import: %d %s", rec.Code, rec.Body)
	}

	if rec := send(r, http.MethodPost, "/api/v1/admin/collections", token, "application/json", strings.NewReader(`{"slug":"nourl"}`)); rec.Code != http.StatusBadRequest {
		t.Errorf("missing url: %d", rec.Code)
	}

	// Production configuration: loopback and private targets are refused.
	guarded := newAdminRouter(fakePinger{}, &fakeCols{}, &fakeOps{}, false)
	token = loginToken(t, guarded)
	rec = send(guarded, http.MethodPost, "/api/v1/admin/collections", token, "application/json", strings.NewReader(`{"url":"`+srv.URL+`/spec.yaml","slug":"internal"}`))
	if rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), "private") {
		t.Errorf("private url: %d %s", rec.Code, rec.Body)
	}
}

func truncate(s string) string {
	if len(s) > 200 {
		return s[:200] + "..."
	}
	return s
}
