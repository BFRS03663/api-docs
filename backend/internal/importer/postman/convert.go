package postman

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// ErrUnsupportedVersion is returned for Postman v1 exports.
var ErrUnsupportedVersion = errors.New("postman collection v1 is not supported; export it as v2.1")

var postmanVar = regexp.MustCompile(`\{\{([^}]+)\}\}`)

// ToOpenAPIJSON converts a Postman v2.1 collection into an OpenAPI 3.0 JSON
// document. The result is not validated here; the caller loads it through the
// OpenAPI parser, which validates and resolves it.
func ToOpenAPIJSON(data []byte) ([]byte, error) {
	var col collection
	if err := json.Unmarshal(data, &col); err != nil {
		return nil, fmt.Errorf("parse collection: %w", err)
	}
	if len(col.Requests) > 0 || (col.Info.Name == "" && len(col.Item) == 0) {
		return nil, ErrUnsupportedVersion
	}

	c := newConverter(col)
	c.walk(col.Item, nil)
	if len(c.paths) == 0 {
		return nil, errors.New("collection contains no requests")
	}
	return json.Marshal(c.document())
}

type converter struct {
	col     collection
	vars    map[string]string
	paths   map[string]map[string]any
	tags    []map[string]any
	tagSeen map[string]bool
	servers []map[string]any
	srvSeen map[string]bool
	schemes map[string]any
	opIDs   map[string]bool
	global  []map[string][]string
}

func newConverter(col collection) *converter {
	c := &converter{
		col:     col,
		vars:    map[string]string{},
		paths:   map[string]map[string]any{},
		tagSeen: map[string]bool{},
		srvSeen: map[string]bool{},
		schemes: map[string]any{},
		opIDs:   map[string]bool{},
	}
	for _, v := range col.Variable {
		if v.Key != "" {
			c.vars[v.Key] = string(v.Value)
		}
	}
	if col.Auth != nil {
		if name := c.registerAuth(col.Auth); name != "" {
			c.global = []map[string][]string{{name: {}}}
		}
	}
	return c
}

// walk descends folders. Every folder becomes a tag whose name is the slug of
// its full path (unique) and whose x-displayName is the folder path joined
// with " / ", so nested Postman folders stay visible as a flat, ordered list.
// (x-tagGroups was tried first, but Scalar then hides every ungrouped tag.)
func (c *converter) walk(items []item, path []string) {
	for _, it := range items {
		if it.Request != nil {
			c.addRequest(it, tagName(path))
			continue
		}
		folderPath := append(append([]string{}, path...), it.Name)
		c.addTag(folderPath, it.Description.Text)
		c.walk(it.Item, folderPath)
	}
}

// tagName is the unique OpenAPI tag for a folder path, "" at the root.
func tagName(path []string) string {
	if len(path) == 0 {
		return ""
	}
	return slugify(strings.Join(path, " "))
}

func (c *converter) addTag(path []string, desc string) {
	name := tagName(path)
	if name == "" || c.tagSeen[name] {
		return
	}
	c.tagSeen[name] = true
	t := map[string]any{"name": name, "x-displayName": strings.Join(path, " / ")}
	if md := toMarkdown(desc); md != "" {
		t["description"] = md
	}
	c.tags = append(c.tags, t)
}

func (c *converter) addRequest(it item, tag string) {
	req := it.Request
	method := strings.ToLower(req.Method)
	if method == "" {
		method = "get"
	}
	server, path, queryFromRaw := c.splitURL(req.URL.Raw)
	c.addServer(server)

	op := map[string]any{
		"summary":     it.Name,
		"operationId": c.uniqueOperationID(it.Name),
		"responses":   c.responses(it.Response),
	}
	if md := toMarkdown(req.Description.Text); md != "" {
		op["description"] = md
	}
	if tag != "" {
		op["tags"] = []string{tag}
	}
	if params := c.parameters(path, req, queryFromRaw); len(params) > 0 {
		op["parameters"] = params
	}
	if rb := c.requestBody(req); rb != nil {
		op["requestBody"] = rb
	}
	switch {
	case req.Auth != nil && req.Auth.Type == "noauth":
		op["security"] = []map[string][]string{}
	case req.Auth != nil:
		if name := c.registerAuth(req.Auth); name != "" {
			op["security"] = []map[string][]string{{name: {}}}
		}
	case bearerHeader(req.Header):
		name := c.registerAuth(&auth{Type: "bearer"})
		op["security"] = []map[string][]string{{name: {}}}
	}

	if c.paths[path] == nil {
		c.paths[path] = map[string]any{}
	}
	if existing, dup := c.paths[path][method].(map[string]any); dup {
		// Postman allows several requests on one endpoint; OpenAPI does not.
		// Fold the later request into the first as named examples so nothing
		// the collection documents is lost.
		mergeOperation(existing, op, it.Name)
		return
	}
	c.paths[path][method] = op
}

// splitURL separates the server from the templated path and any query string.
// A leading {{var}} is the server: its value when the collection defines it
// (Postman's baseUrl convention, which may include a path prefix), otherwise
// an OpenAPI server variable. Absolute URLs split at the host.
func (c *converter) splitURL(raw string) (server, path, query string) {
	raw = strings.TrimSpace(raw)
	if i := strings.IndexByte(raw, '?'); i >= 0 {
		raw, query = raw[:i], raw[i+1:]
	}
	if i := strings.IndexByte(raw, '#'); i >= 0 {
		raw = raw[:i]
	}

	if m := postmanVar.FindStringSubmatchIndex(raw); m != nil && m[0] == 0 {
		name := raw[m[2]:m[3]]
		if v, ok := c.vars[name]; ok {
			server = strings.TrimSuffix(strings.TrimSpace(v), "/")
		} else {
			server = "{" + name + "}"
		}
		return server, templatePath(c.substitute(raw[m[1]:])), query
	}

	raw = c.substitute(raw)
	if strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://") {
		rest := raw[strings.Index(raw, "://")+3:]
		if i := strings.IndexByte(rest, '/'); i >= 0 {
			server, path = raw[:len(raw)-len(rest)+i], rest[i:]
		} else {
			server, path = raw, "/"
		}
		return server, templatePath(path), query
	}
	return "", templatePath(raw), query
}

// templatePath turns Postman's :id and {{id}} segments into OpenAPI {id}.
func templatePath(p string) string {
	p = strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' || r == '\t' || r == ' ' {
			return -1
		}
		return r
	}, p)
	p = postmanVar.ReplaceAllString(p, "{$1}")
	segs := strings.Split(p, "/")
	for i, s := range segs {
		if strings.HasPrefix(s, ":") && len(s) > 1 {
			segs[i] = "{" + s[1:] + "}"
		}
	}
	p = strings.Join(segs, "/")
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	if len(p) > 1 {
		p = strings.TrimSuffix(p, "/")
	}
	return p
}

func (c *converter) substitute(s string) string {
	return postmanVar.ReplaceAllStringFunc(s, func(m string) string {
		name := m[2 : len(m)-2]
		if v, ok := c.vars[name]; ok {
			return v
		}
		return m
	})
}

func (c *converter) addServer(server string) {
	if server == "" || c.srvSeen[server] {
		return
	}
	c.srvSeen[server] = true
	entry := map[string]any{"url": server}
	if strings.HasPrefix(server, "{") {
		name := strings.Trim(server, "{}")
		entry["variables"] = map[string]any{name: map[string]any{
			"default":     "http://localhost",
			"description": "Postman variable " + name + " was not defined in the collection",
		}}
	}
	c.servers = append(c.servers, entry)
}

func (c *converter) uniqueOperationID(name string) string {
	base := slugify(name)
	if base == "" {
		base = "operation"
	}
	id := base
	for n := 2; c.opIDs[id]; n++ {
		id = fmt.Sprintf("%s-%d", base, n)
	}
	c.opIDs[id] = true
	return id
}

var nonSlug = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(s string) string {
	return strings.Trim(nonSlug.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

// registerAuth maps a Postman auth block to a security scheme and returns
// its name, or "" for types that cannot be represented.
func (c *converter) registerAuth(a *auth) string {
	switch a.Type {
	case "bearer", "oauth2", "jwt":
		c.schemes["bearerAuth"] = map[string]any{"type": "http", "scheme": "bearer"}
		return "bearerAuth"
	case "basic":
		c.schemes["basicAuth"] = map[string]any{"type": "http", "scheme": "basic"}
		return "basicAuth"
	case "apikey":
		name, in := "X-API-Key", "header"
		for _, p := range a.APIKey {
			switch p.Key {
			case "key":
				if p.Value != "" {
					name = string(p.Value)
				}
			case "in":
				if p.Value == "query" {
					in = "query"
				}
			}
		}
		c.schemes["apiKeyAuth"] = map[string]any{"type": "apiKey", "in": in, "name": name}
		return "apiKeyAuth"
	}
	return ""
}

func (c *converter) document() map[string]any {
	version := "1.0.0"
	if v, ok := c.vars["version"]; ok && v != "" {
		version = v
	}
	info := map[string]any{"title": c.col.Info.Name, "version": version}
	if md := toMarkdown(c.col.Info.Description.Text); md != "" {
		info["description"] = md
	}
	doc := map[string]any{
		"openapi": "3.0.3",
		"info":    info,
		"paths":   c.paths,
	}
	if len(c.tags) > 0 {
		doc["tags"] = c.tags
	}
	if len(c.servers) > 0 {
		doc["servers"] = c.servers
	}
	if len(c.schemes) > 0 {
		doc["components"] = map[string]any{"securitySchemes": c.schemes}
	}
	if len(c.global) > 0 {
		doc["security"] = c.global
	}
	return doc
}

// mergeOperation folds op into base: descriptions are appended under a
// heading, request bodies and responses become multi-example media types,
// and parameters are unioned by (in, name).
func mergeOperation(base, op map[string]any, name string) {
	baseTags, _ := base["tags"].([]string)
	if extra, ok := op["tags"].([]string); ok {
		for _, t := range extra {
			if !slices.Contains(baseTags, t) {
				baseTags = append(baseTags, t)
			}
		}
		base["tags"] = baseTags
	}
	if d, _ := op["description"].(string); d != "" {
		prev, _ := base["description"].(string)
		if !strings.Contains(prev, d) {
			base["description"] = strings.TrimSpace(prev + "\n\n### " + name + "\n\n" + d)
		}
	}
	baseName, _ := base["summary"].(string)

	if rb, ok := op["requestBody"].(map[string]any); ok {
		if baseRB, ok := base["requestBody"].(map[string]any); ok {
			mergeContent(baseRB["content"].(map[string]any), rb["content"].(map[string]any), baseName, name)
		} else {
			promoteAll(rb["content"].(map[string]any), name)
			base["requestBody"] = rb
		}
	}

	baseResp := base["responses"].(map[string]any)
	for code, r := range op["responses"].(map[string]any) {
		resp := r.(map[string]any)
		existing, ok := baseResp[code].(map[string]any)
		if !ok {
			if content, ok := resp["content"].(map[string]any); ok {
				promoteAll(content, name)
			}
			baseResp[code] = resp
			continue
		}
		ec, _ := existing["content"].(map[string]any)
		nc, _ := resp["content"].(map[string]any)
		if ec == nil && nc != nil {
			promoteAll(nc, name)
			existing["content"] = nc
		} else if ec != nil && nc != nil {
			mergeContent(ec, nc, baseName, name)
		}
	}

	seen := map[string]bool{}
	baseParams, _ := base["parameters"].([]map[string]any)
	for _, p := range baseParams {
		seen[p["in"].(string)+":"+p["name"].(string)] = true
	}
	if extra, ok := op["parameters"].([]map[string]any); ok {
		for _, p := range extra {
			if key := p["in"].(string) + ":" + p["name"].(string); !seen[key] {
				seen[key] = true
				baseParams = append(baseParams, p)
			}
		}
		if len(baseParams) > 0 {
			base["parameters"] = baseParams
		}
	}
}

// mergeContent adds every media type in src to dst, converting single
// "example" values into named "examples" so both variants stay visible.
func mergeContent(dst, src map[string]any, dstName, srcName string) {
	for ct, m := range src {
		media := m.(map[string]any)
		target, ok := dst[ct].(map[string]any)
		if !ok {
			promote(media, srcName)
			dst[ct] = media
			continue
		}
		promote(target, dstName)
		promote(media, srcName)
		examples := target["examples"].(map[string]any)
		for k, v := range media["examples"].(map[string]any) {
			if hasEqualExample(examples, v) {
				continue // same payload documented twice; keep one
			}
			key := k
			for n := 2; ; n++ {
				if _, taken := examples[key]; !taken {
					break
				}
				key = fmt.Sprintf("%s-%d", k, n)
			}
			examples[key] = v
		}
	}
}

// hasEqualExample reports whether an example with the same value already exists.
func hasEqualExample(examples map[string]any, candidate any) bool {
	want, err := json.Marshal(candidate.(map[string]any)["value"])
	if err != nil {
		return false
	}
	for _, ex := range examples {
		got, err := json.Marshal(ex.(map[string]any)["value"])
		if err == nil && string(got) == string(want) {
			return true
		}
	}
	return false
}

func promoteAll(content map[string]any, name string) {
	for _, m := range content {
		promote(m.(map[string]any), name)
	}
}

// promote rewrites {"example": v} as {"examples": {slug(name): {summary, value}}}.
func promote(media map[string]any, name string) {
	if _, done := media["examples"]; done {
		return
	}
	examples := map[string]any{}
	if v, ok := media["example"]; ok {
		key := slugify(name)
		if key == "" {
			key = "example"
		}
		examples[key] = map[string]any{"summary": name, "value": v}
		delete(media, "example")
	}
	media["examples"] = examples
}
