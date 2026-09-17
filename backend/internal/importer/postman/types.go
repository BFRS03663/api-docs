// Package postman converts Postman Collection v2.1 exports into OpenAPI 3.0
// JSON so the rest of the system only ever deals with OpenAPI.
package postman

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// collection mirrors the subset of the v2.1 schema we need.
type collection struct {
	Info     info            `json:"info"`
	Item     []item          `json:"item"`
	Auth     *auth           `json:"auth"`
	Variable []variable      `json:"variable"`
	Requests json.RawMessage `json:"requests"` // present only in v1 exports
}

type info struct {
	Name        string      `json:"name"`
	Description description `json:"description"`
	Schema      string      `json:"schema"`
	PostmanID   string      `json:"_postman_id"`
}

type item struct {
	Name        string      `json:"name"`
	Description description `json:"description"`
	Item        []item      `json:"item"`
	Request     *request    `json:"request"`
	Response    []response  `json:"response"`
}

type request struct {
	Method      string      `json:"method"`
	Header      []kv        `json:"header"`
	URL         requestURL  `json:"url"`
	URLObject   *requestURL `json:"urlObject"` // Postman v2.0 exports put query/variable here when url is a string
	Body        *body       `json:"body"`
	Description description `json:"description"`
	Auth        *auth       `json:"auth"`
}

// UnmarshalJSON accepts the shorthand where a request is just a URL string.
func (r *request) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var raw string
		if err := json.Unmarshal(data, &raw); err != nil {
			return err
		}
		*r = request{Method: "GET", URL: requestURL{Raw: raw}}
		return nil
	}
	type plain request
	var p plain
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	*r = request(p)
	if r.URLObject != nil {
		if len(r.URL.Query) == 0 {
			r.URL.Query = r.URLObject.Query
		}
		if len(r.URL.Variable) == 0 {
			r.URL.Variable = r.URLObject.Variable
		}
	}
	return nil
}

type requestURL struct {
	Raw      string   `json:"raw"`
	Host     []string `json:"host"`
	Path     []string `json:"path"`
	Query    []kv     `json:"query"`
	Variable []kv     `json:"variable"`
}

// UnmarshalJSON accepts either a URL string or the structured object. Path
// segments that are objects ({value,type}) are reduced to their value.
func (u *requestURL) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		return json.Unmarshal(data, &u.Raw)
	}
	var p struct {
		Raw      string            `json:"raw"`
		Host     flexStrings       `json:"host"`
		Path     []json.RawMessage `json:"path"`
		Query    []kv              `json:"query"`
		Variable []kv              `json:"variable"`
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	u.Raw, u.Host, u.Query, u.Variable = p.Raw, p.Host, p.Query, p.Variable
	for _, seg := range p.Path {
		var s string
		if err := json.Unmarshal(seg, &s); err == nil {
			u.Path = append(u.Path, s)
			continue
		}
		var obj struct {
			Value string `json:"value"`
		}
		if err := json.Unmarshal(seg, &obj); err == nil {
			u.Path = append(u.Path, obj.Value)
		}
	}
	if u.Raw == "" && (len(u.Host) > 0 || len(u.Path) > 0) {
		u.Raw = strings.Join(u.Host, ".") + "/" + strings.Join(u.Path, "/")
	}
	return nil
}

type body struct {
	Mode       string `json:"mode"`
	Raw        string `json:"raw"`
	URLEncoded []kv   `json:"urlencoded"`
	FormData   []kv   `json:"formdata"`
	Options    struct {
		Raw struct {
			Language string `json:"language"`
		} `json:"raw"`
	} `json:"options"`
}

type auth struct {
	Type   string     `json:"type"`
	Bearer authParams `json:"bearer"`
	APIKey authParams `json:"apikey"`
	Basic  authParams `json:"basic"`
}

// authParams accepts the v2.1 list form ([{key,value}]) and the v2.0 object
// form ({key: value}) of an auth block's parameters.
type authParams []kv

func (a *authParams) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" || trimmed == "null" {
		return nil
	}
	if trimmed[0] == '[' {
		var list []kv
		if err := json.Unmarshal(data, &list); err != nil {
			return err
		}
		*a = list
		return nil
	}
	var m map[string]flexString
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	for k, v := range m {
		*a = append(*a, kv{Key: k, Value: v})
	}
	return nil
}

type response struct {
	Name            string  `json:"name"`
	Code            flexInt `json:"code"`
	Status          string  `json:"status"`
	Header          []kv    `json:"header"`
	Body            string  `json:"body"`
	PreviewLanguage string  `json:"_postman_previewlanguage"`
}

// flexInt accepts a JSON number, a numeric string, or null.
type flexInt int

func (f *flexInt) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "null" {
		*f = 0
		return nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("response code %q is not a number", s)
	}
	*f = flexInt(n)
	return nil
}

type variable struct {
	Key   string     `json:"key"`
	Value flexString `json:"value"`
}

type kv struct {
	Key         string      `json:"key"`
	Value       flexString  `json:"value"`
	Description description `json:"description"`
	Disabled    bool        `json:"disabled"`
	Type        string      `json:"type"`
}

// description is either a plain string or {"content": "...", "type": "..."}.
type description struct {
	Text string
}

func (d *description) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	if data[0] == '"' {
		return json.Unmarshal(data, &d.Text)
	}
	var obj struct {
		Content string `json:"content"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	d.Text = obj.Content
	return nil
}

// flexString accepts strings, numbers, booleans and null.
type flexString string

func (f *flexString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*f = flexString(s)
		return nil
	}
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*f = flexString(fmt.Sprint(v))
	return nil
}

// flexStrings accepts either a string or an array of strings (Postman's host).
type flexStrings []string

func (f *flexStrings) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*f = flexStrings{s}
		return nil
	}
	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	*f = arr
	return nil
}
