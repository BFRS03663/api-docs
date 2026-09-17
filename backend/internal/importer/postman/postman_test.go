package postman

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func convertFixture(t *testing.T) map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "..", "testdata", "sample.postman_collection.json"))
	if err != nil {
		t.Fatal(err)
	}
	js, err := ToOpenAPIJSON(data)
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	var doc map[string]any
	if err := json.Unmarshal(js, &doc); err != nil {
		t.Fatal(err)
	}
	return doc
}

func dig(t *testing.T, m any, keys ...string) any {
	t.Helper()
	cur := m
	for _, k := range keys {
		obj, ok := cur.(map[string]any)
		if !ok {
			t.Fatalf("expected object at %q, got %T", k, cur)
		}
		cur, ok = obj[k]
		if !ok {
			t.Fatalf("key %q missing (path %v)", k, keys)
		}
	}
	return cur
}

func TestConvertFixture(t *testing.T) {
	doc := convertFixture(t)

	if dig(t, doc, "openapi") != "3.0.3" || dig(t, doc, "info", "title") != "Acme Users API" {
		t.Fatalf("header wrong: %v", doc["info"])
	}

	paths := doc["paths"].(map[string]any)
	want := []string{"/users", "/users/{id}", "/users/{userId}/active", "/users/{id}/avatar", "/login"}
	for _, p := range want {
		if _, ok := paths[p]; !ok {
			t.Errorf("path %s missing; have %v", p, keys(paths))
		}
	}

	// Variables substituted into the server; the auth host becomes a second server.
	servers := doc["servers"].([]any)
	if len(servers) != 2 || dig(t, servers[0], "url") != "https://api.example.com/v1" || dig(t, servers[1], "url") != "https://auth.example.com" {
		t.Errorf("servers = %v", servers)
	}

	// Folder → tag, path variable → path param with description, query params kept, disabled dropped.
	get := dig(t, paths, "/users/{id}", "get").(map[string]any)
	if tags := get["tags"].([]any); len(tags) != 1 || tags[0] != "users" {
		t.Errorf("tags = %v", tags)
	}
	params := get["parameters"].([]any)
	if len(params) != 1 || dig(t, params[0], "name") != "id" || dig(t, params[0], "in") != "path" || dig(t, params[0], "required") != true || dig(t, params[0], "description") != "User identifier" {
		t.Errorf("path params = %v", params)
	}
	list := dig(t, paths, "/users", "get").(map[string]any)
	lp := list["parameters"].([]any)
	names := map[string]bool{}
	for _, p := range lp {
		names[dig(t, p, "name").(string)] = true
	}
	if !names["page"] || !names["limit"] || !names["Accept"] || names["debug"] {
		t.Errorf("list params = %v", names)
	}

	// Saved responses → responses with parsed JSON examples and inferred schema.
	resp := get["responses"].(map[string]any)
	if _, ok := resp["404"]; !ok {
		t.Errorf("404 response missing: %v", keys(resp))
	}
	ex := dig(t, resp, "200", "content", "application/json", "example").(map[string]any)
	if ex["id"] != float64(42) {
		t.Errorf("200 example = %v", ex)
	}
	if dig(t, resp, "200", "content", "application/json", "schema", "properties", "name", "type") != "string" {
		t.Error("schema not inferred from example")
	}

	// Raw JSON body → requestBody; markdown description object flattened.
	post := dig(t, paths, "/users", "post").(map[string]any)
	if post["description"] != "Creates a user." {
		t.Errorf("description = %v", post["description"])
	}
	body := dig(t, post, "requestBody", "content", "application/json", "example").(map[string]any)
	if body["email"] != "ada@example.com" {
		t.Errorf("request example = %v", body)
	}
	if _, has := post["security"]; has {
		t.Error("request without auth should inherit global security, not override it")
	}

	// Nested folder keeps parent tag; {{userId}} became a path param.
	del := dig(t, paths, "/users/{userId}/active", "delete").(map[string]any)
	if tags := del["tags"].([]any); tags[0] != "users-admin" {
		t.Errorf("nested tag = %v", tags)
	}

	// formdata → multipart with binary file; noauth → empty security.
	put := dig(t, paths, "/users/{id}/avatar", "put").(map[string]any)
	if dig(t, put, "requestBody", "content", "multipart/form-data", "schema", "properties", "file", "format") != "binary" {
		t.Error("file field not binary")
	}
	if sec, ok := put["security"].([]any); !ok || len(sec) != 0 {
		t.Errorf("noauth security = %v", put["security"])
	}
	if put["responses"].(map[string]any)["200"] == nil {
		t.Error("default 200 response missing")
	}

	// urlencoded → form schema.
	login := dig(t, paths, "/login", "post").(map[string]any)
	if dig(t, login, "requestBody", "content", "application/x-www-form-urlencoded", "schema", "properties", "username", "type") != "string" {
		t.Error("urlencoded schema missing")
	}

	// Collection auth → bearer scheme + global security.
	if dig(t, doc, "components", "securitySchemes", "bearerAuth", "scheme") != "bearer" {
		t.Error("bearer scheme missing")
	}
	if sec := doc["security"].([]any); len(sec) != 1 {
		t.Errorf("global security = %v", sec)
	}
	tags := doc["tags"].([]any)
	if len(tags) != 2 || dig(t, tags[0], "description") != "Manage user accounts." || dig(t, tags[0], "x-displayName") != "Users" || dig(t, tags[1], "x-displayName") != "Users / Admin" {
		t.Errorf("tags = %v", tags)
	}
	if _, has := doc["x-tagGroups"]; has {
		t.Error("tag groups must not be emitted: Scalar hides ungrouped tags")
	}
}

// TestV20DocumenterShapes covers what Postman's published-docs export emits:
// string response codes, urlObject beside a string url, HTML descriptions,
// Authorization headers instead of auth blocks, and preview languages.
func TestV20DocumenterShapes(t *testing.T) {
	src := `{"info":{"_postman_id":"x","name":"Docs","description":"<html><body><h1>Intro</h1><p>Hello <strong>world</strong></p><table><thead><tr><th>Field</th><th>Type</th></tr></thead><tbody><tr><td>id</td><td>int</td></tr></tbody></table></body></html>","schema":"https://schema.getpostman.com/json/collection/v2.0.0/collection.json"},
	"item":[{"name":"Orders","description":"<p>Order APIs</p>","item":[
	  {"name":"List","request":{"method":"GET","header":[{"key":"Authorization","value":"Bearer {{token}}","type":"text"}],"url":"https://api.example.com/v1/orders?page=1","urlObject":{"query":[{"key":"page","value":"1","description":"Page"}],"variable":[]},"description":"<p>Lists orders.</p>"},
	   "response":[{"name":"OK","code":"200","status":"OK","_postman_previewlanguage":"json","header":[],"body":"{\"data\":[]}"},{"name":"Err","code":null,"status":"","header":null,"body":"oops"}]}
	]}]}`
	js, err := ToOpenAPIJSON([]byte(src))
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	var doc map[string]any
	if err := json.Unmarshal(js, &doc); err != nil {
		t.Fatal(err)
	}
	desc := dig(t, doc, "info", "description").(string)
	if !strings.Contains(desc, "# Intro") || !strings.Contains(desc, "**world**") || !strings.Contains(desc, "| Field |") || strings.Contains(desc, "<p>") {
		t.Errorf("html not converted to markdown:\n%s", desc)
	}
	get := dig(t, doc, "paths", "/v1/orders", "get").(map[string]any)
	if get["description"] != "Lists orders." {
		t.Errorf("request description = %q", get["description"])
	}
	params := get["parameters"].([]any)
	if len(params) != 1 || dig(t, params[0], "name") != "page" || dig(t, params[0], "description") != "Page" {
		t.Errorf("urlObject query not used: %v", params)
	}
	if sec := get["security"].([]any); len(sec) != 1 {
		t.Errorf("Authorization header should imply bearer security: %v", get["security"])
	}
	if dig(t, doc, "components", "securitySchemes", "bearerAuth", "scheme") != "bearer" {
		t.Error("bearer scheme missing")
	}
	resp := get["responses"].(map[string]any)
	if _, ok := resp["200"]; !ok {
		t.Errorf("string code not parsed: %v", keys(resp))
	}
	if dig(t, resp, "200", "content", "application/json", "example").(map[string]any)["data"] == nil {
		t.Error("preview language json not honoured")
	}
	if dig(t, doc, "tags").([]any)[0].(map[string]any)["description"] != "Order APIs" {
		t.Error("folder html description not converted")
	}
}

func TestUnsupportedAndEmpty(t *testing.T) {
	_, err := ToOpenAPIJSON([]byte(`{"id":"x","name":"old","requests":[{"url":"http://a/b"}]}`))
	if !errors.Is(err, ErrUnsupportedVersion) {
		t.Fatalf("v1 err = %v", err)
	}
	_, err = ToOpenAPIJSON([]byte(`{"info":{"name":"empty"},"item":[{"name":"folder","item":[]}]}`))
	if err == nil {
		t.Fatal("expected error for collection with no requests")
	}
	_, err = ToOpenAPIJSON([]byte(`not json`))
	if err == nil {
		t.Fatal("expected parse error")
	}
}

func TestDuplicateEndpointsMergeIntoExamples(t *testing.T) {
	src := `{"info":{"_postman_id":"x","name":"Dup"},"item":[
	 {"name":"Create order","request":{"method":"POST","header":[{"key":"Content-Type","value":"application/json"}],"url":"https://a.io/orders","body":{"mode":"raw","raw":"{\"kind\":\"basic\"}"},"description":"Basic."},
	  "response":[{"name":"OK","code":200,"header":[{"key":"Content-Type","value":"application/json"}],"body":"{\"id\":1}"}]},
	 {"name":"Same again","request":{"method":"POST","header":[{"key":"Content-Type","value":"application/json"}],"url":"https://a.io/orders","body":{"mode":"raw","raw":"{\"kind\":\"basic\"}"},"description":"Basic."},"response":[]},
	 {"name":"Create order (advanced)","request":{"method":"POST","header":[{"key":"Content-Type","value":"application/json"}],"url":"https://a.io/orders?dry=1","urlObject":{"query":[{"key":"dry","value":"1"}]},"body":{"mode":"raw","raw":"{\"kind\":\"advanced\"}"},"description":"Advanced."},
	  "response":[{"name":"OK","code":200,"header":[{"key":"Content-Type","value":"application/json"}],"body":"{\"id\":2}"},{"name":"Bad","code":422,"body":"{\"error\":\"x\"}"}]}
	]}`
	js, err := ToOpenAPIJSON([]byte(src))
	if err != nil {
		t.Fatalf("convert: %v", err)
	}
	var doc map[string]any
	_ = json.Unmarshal(js, &doc)
	post := dig(t, doc, "paths", "/orders", "post").(map[string]any)
	if post["summary"] != "Create order" || !strings.Contains(post["description"].(string), "### Create order (advanced)") {
		t.Errorf("summary/description = %v / %v", post["summary"], post["description"])
	}
	reqEx := dig(t, post, "requestBody", "content", "application/json", "examples").(map[string]any)
	if len(reqEx) != 2 || dig(t, reqEx, "create-order-advanced", "value", "kind") != "advanced" {
		t.Errorf("identical example should be deduped, distinct kept: %v", reqEx)
	}
	if strings.Count(post["description"].(string), "Basic.") != 1 {
		t.Errorf("identical description repeated: %q", post["description"])
	}
	if _, single := dig(t, post, "requestBody", "content", "application/json").(map[string]any)["example"]; single {
		t.Error("single example should have been promoted")
	}
	respEx := dig(t, post, "responses", "200", "content", "application/json", "examples").(map[string]any)
	if len(respEx) != 2 {
		t.Errorf("200 examples = %v", respEx)
	}
	if _, ok := dig(t, post, "responses").(map[string]any)["422"]; !ok {
		t.Error("422 from second request missing")
	}
	params := post["parameters"].([]any)
	if len(params) != 1 || dig(t, params[0], "name") != "dry" {
		t.Errorf("params not unioned: %v", params)
	}
}

func TestSameEndpointInTwoFoldersGetsBothTags(t *testing.T) {
	src := `{"info":{"_postman_id":"x","name":"T"},"item":[
	 {"name":"Orders","item":[{"name":"List","request":{"method":"GET","url":"https://a.io/orders"},"response":[]}]},
	 {"name":"Hyperlocal","item":[{"name":"Orders","item":[{"name":"List","request":{"method":"GET","url":"https://a.io/orders"},"response":[]}]}]}
	]}`
	js, err := ToOpenAPIJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string]any
	_ = json.Unmarshal(js, &doc)
	tags := dig(t, doc, "paths", "/orders", "get", "tags").([]any)
	if len(tags) != 2 || tags[0] != "orders" || tags[1] != "hyperlocal-orders" {
		t.Errorf("tags = %v", tags)
	}
}

func TestSameStatusCodeResponsesBecomeExamples(t *testing.T) {
	src := `{"info":{"_postman_id":"x","name":"R"},"item":[{"name":"Create","request":{"method":"POST","url":"https://a.io/x"},
	 "response":[
	  {"name":"Successful Call","code":200,"header":[{"key":"Content-Type","value":"application/json"}],"body":"{\"ok\":true}"},
	  {"name":"Invalid Data","code":400,"header":[{"key":"Content-Type","value":"application/json"}],"body":"{\"error\":\"invalid\"}"},
	  {"name":"Missing Fields","code":400,"header":[{"key":"Content-Type","value":"application/json"}],"body":"{\"error\":\"missing\"}"}
	 ]}]}`
	js, err := ToOpenAPIJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string]any
	_ = json.Unmarshal(js, &doc)
	r400 := dig(t, doc, "paths", "/x", "post", "responses", "400").(map[string]any)
	if r400["description"] != "Invalid Data" {
		t.Errorf("description = %v", r400["description"])
	}
	ex := dig(t, r400, "content", "application/json", "examples").(map[string]any)
	if len(ex) != 2 || dig(t, ex, "missing-fields", "value", "error") != "missing" || dig(t, ex, "invalid-data", "value", "error") != "invalid" {
		t.Errorf("400 examples = %v", ex)
	}
	if _, single := dig(t, doc, "paths", "/x", "post", "responses", "200", "content", "application/json").(map[string]any)["example"]; !single {
		t.Error("lone 200 response should keep its plain example")
	}
}

func TestAuthParamsBothShapes(t *testing.T) {
	var list, obj auth
	if err := json.Unmarshal([]byte(`{"type":"apikey","apikey":[{"key":"key","value":"X-Token"},{"key":"in","value":"query"}]}`), &list); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(`{"type":"apikey","apikey":{"key":"X-Token","in":"query"}}`), &obj); err != nil {
		t.Fatal(err)
	}
	for _, a := range []auth{list, obj} {
		c := newConverter(collection{})
		if name := c.registerAuth(&a); name != "apiKeyAuth" {
			t.Fatalf("name = %q", name)
		}
		scheme := c.schemes["apiKeyAuth"].(map[string]any)
		if scheme["name"] != "X-Token" || scheme["in"] != "query" {
			t.Errorf("scheme = %v", scheme)
		}
	}
}

func TestTemplatePathAndSplit(t *testing.T) {
	cases := map[string]string{
		"users/:id/posts/:postId": "/users/{id}/posts/{postId}",
		"/track/awb/{{awb}}\n":    "/track/awb/{awb}",
		"/a/{{b}}/c/":             "/a/{b}/c",
		"":                        "/",
	}
	for in, want := range cases {
		if got := templatePath(in); got != want {
			t.Errorf("templatePath(%q) = %q, want %q", in, got, want)
		}
	}
	c := newConverter(collection{Variable: []variable{{Key: "host", Value: "https://x.io"}}})
	server, path, query := c.splitURL("{{host}}/v2/things?x=1&y=2")
	if server != "https://x.io" || path != "/v2/things" || query != "x=1&y=2" {
		t.Errorf("split = %q %q %q", server, path, query)
	}
	server, path, _ = c.splitURL("{{unknown}}/things")
	if server != "{unknown}" || path != "/things" {
		t.Errorf("unknown var split = %q %q", server, path)
	}
	if id1, id2 := c.uniqueOperationID("Get User"), c.uniqueOperationID("Get user!"); id1 != "get-user" || id2 != "get-user-2" {
		t.Errorf("ids = %q %q", id1, id2)
	}
}

func keys(m map[string]any) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
