package postman

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var pathParam = regexp.MustCompile(`\{([^}/]+)\}`)

// parameters builds path, query and header parameters for one request.
func (c *converter) parameters(path string, req *request, rawQuery string) []map[string]any {
	var params []map[string]any
	seen := map[string]bool{}

	descByKey := map[string]string{}
	for _, v := range req.URL.Variable {
		descByKey[v.Key] = v.Description.Text
	}
	for _, m := range pathParam.FindAllStringSubmatch(path, -1) {
		name := m[1]
		if seen["path:"+name] {
			continue
		}
		seen["path:"+name] = true
		params = append(params, param(name, "path", true, descByKey[name], ""))
	}

	queries := req.URL.Query
	if len(queries) == 0 && rawQuery != "" {
		if values, err := url.ParseQuery(rawQuery); err == nil {
			for k, vs := range values {
				queries = append(queries, kv{Key: k, Value: flexString(strings.Join(vs, ","))})
			}
		}
	}
	for _, q := range queries {
		if q.Disabled || q.Key == "" || seen["query:"+q.Key] {
			continue
		}
		seen["query:"+q.Key] = true
		params = append(params, param(q.Key, "query", false, q.Description.Text, string(q.Value)))
	}

	for _, h := range req.Header {
		key := strings.ToLower(h.Key)
		if h.Disabled || h.Key == "" || key == "content-type" || key == "authorization" || seen["header:"+key] {
			continue
		}
		seen["header:"+key] = true
		params = append(params, param(h.Key, "header", false, h.Description.Text, string(h.Value)))
	}
	return params
}

func param(name, in string, required bool, desc, example string) map[string]any {
	p := map[string]any{
		"name":   name,
		"in":     in,
		"schema": map[string]any{"type": "string"},
	}
	if required {
		p["required"] = true
	}
	if desc != "" {
		p["description"] = desc
	}
	if example != "" {
		p["example"] = example
	}
	return p
}

// requestBody maps Postman body modes to an OpenAPI requestBody.
func (c *converter) requestBody(req *request) map[string]any {
	b := req.Body
	if b == nil {
		return nil
	}
	switch b.Mode {
	case "raw":
		if strings.TrimSpace(b.Raw) == "" {
			return nil
		}
		ct := headerValue(req.Header, "Content-Type")
		if ct == "" {
			ct = contentTypeForLanguage(b.Options.Raw.Language)
		}
		media := map[string]any{"schema": schemaFor(b.Raw, ct)}
		media["example"] = exampleFor(b.Raw, ct)
		return map[string]any{"content": map[string]any{ct: media}}
	case "urlencoded":
		return formBody("application/x-www-form-urlencoded", b.URLEncoded)
	case "formdata":
		return formBody("multipart/form-data", b.FormData)
	}
	return nil
}

func formBody(ct string, fields []kv) map[string]any {
	props := map[string]any{}
	for _, f := range fields {
		if f.Disabled || f.Key == "" {
			continue
		}
		prop := map[string]any{"type": "string"}
		if f.Type == "file" {
			prop["format"] = "binary"
		} else if f.Value != "" {
			prop["example"] = string(f.Value)
		}
		if f.Description.Text != "" {
			prop["description"] = f.Description.Text
		}
		props[f.Key] = prop
	}
	if len(props) == 0 {
		return nil
	}
	schema := map[string]any{"type": "object", "properties": props}
	return map[string]any{"content": map[string]any{ct: map[string]any{"schema": schema}}}
}

// responses converts saved example responses; OpenAPI requires at least one.
// Several saved responses with the same status code become named examples of
// that response so every documented case stays visible.
func (c *converter) responses(saved []response) map[string]any {
	out := map[string]any{}
	for _, r := range saved {
		code := strconv.Itoa(int(r.Code))
		if r.Code == 0 {
			code = "200"
		}
		name := r.Name
		if name == "" {
			name = r.Status
		}
		if name == "" {
			name = "Response"
		}
		var content map[string]any
		if strings.TrimSpace(r.Body) != "" {
			ct := headerValue(r.Header, "Content-Type")
			if ct == "" {
				ct = contentTypeForLanguage(r.PreviewLanguage)
				if ct == "text/plain" && json.Valid([]byte(r.Body)) {
					ct = "application/json"
				}
			}
			content = map[string]any{ct: map[string]any{
				"schema":  schemaFor(r.Body, ct),
				"example": exampleFor(r.Body, ct),
			}}
		}

		existing, dup := out[code].(map[string]any)
		if !dup {
			entry := map[string]any{"description": name}
			if content != nil {
				entry["content"] = content
			}
			out[code] = entry
			continue
		}
		// Same code again: keep the first description, add this body as an example.
		if content == nil {
			continue
		}
		ec, ok := existing["content"].(map[string]any)
		if !ok {
			promoteAll(content, name)
			existing["content"] = content
			continue
		}
		mergeContent(ec, content, existing["description"].(string), name)
	}
	if len(out) == 0 {
		out["200"] = map[string]any{"description": "Successful response"}
	}
	return out
}

func headerValue(headers []kv, name string) string {
	for _, h := range headers {
		if !h.Disabled && strings.EqualFold(h.Key, name) {
			v := string(h.Value)
			if i := strings.IndexByte(v, ';'); i >= 0 {
				v = v[:i]
			}
			return strings.TrimSpace(v)
		}
	}
	return ""
}

// bearerHeader reports whether the request carries an Authorization header,
// which Postman v2.0 exports use instead of an auth block.
func bearerHeader(headers []kv) bool {
	for _, h := range headers {
		if !h.Disabled && strings.EqualFold(h.Key, "Authorization") {
			return true
		}
	}
	return false
}

func contentTypeForLanguage(lang string) string {
	switch strings.ToLower(lang) {
	case "json":
		return "application/json"
	case "xml":
		return "application/xml"
	case "html":
		return "text/html"
	case "javascript":
		return "application/javascript"
	}
	return "text/plain"
}

func isJSON(ct string) bool {
	return strings.Contains(ct, "json")
}

// exampleFor returns parsed JSON for JSON media types (so the renderer can
// pretty-print it) and the raw string otherwise.
func exampleFor(raw, ct string) any {
	if isJSON(ct) {
		var v any
		if err := json.Unmarshal([]byte(raw), &v); err == nil {
			return v
		}
	}
	return raw
}

// schemaFor infers a shallow schema from a JSON example so the reference
// shows field names and types instead of an opaque blob.
func schemaFor(raw, ct string) map[string]any {
	if isJSON(ct) {
		var v any
		if err := json.Unmarshal([]byte(raw), &v); err == nil {
			return inferSchema(v, 0)
		}
	}
	return map[string]any{"type": "string"}
}

func inferSchema(v any, depth int) map[string]any {
	switch t := v.(type) {
	case map[string]any:
		props := map[string]any{}
		if depth < 3 {
			for k, val := range t {
				props[k] = inferSchema(val, depth+1)
			}
		}
		return map[string]any{"type": "object", "properties": props}
	case []any:
		items := map[string]any{}
		if len(t) > 0 && depth < 3 {
			items = inferSchema(t[0], depth+1)
		}
		return map[string]any{"type": "array", "items": items}
	case string:
		return map[string]any{"type": "string"}
	case bool:
		return map[string]any{"type": "boolean"}
	case float64:
		if t == float64(int64(t)) {
			return map[string]any{"type": "integer"}
		}
		return map[string]any{"type": "number"}
	}
	return map[string]any{}
}
