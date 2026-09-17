package markdown

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

const maxSchemaDepth = 3

// writeSchema renders a schema as an indented bullet list of fields.
func writeSchema(b *strings.Builder, ref *openapi3.SchemaRef) {
	if ref == nil || ref.Value == nil {
		return
	}
	v := ref.Value
	switch {
	case len(v.Properties) > 0:
		b.WriteString("Fields:\n\n")
		writeProperties(b, v, 0)
		b.WriteString("\n")
	case v.Items != nil && v.Items.Value != nil && len(v.Items.Value.Properties) > 0:
		b.WriteString("Array of objects with fields:\n\n")
		writeProperties(b, v.Items.Value, 0)
		b.WriteString("\n")
	default:
		if t := typeName(ref); t != "" && t != "object" {
			fmt.Fprintf(b, "Type: %s\n\n", t)
		}
	}
}

func writeProperties(b *strings.Builder, v *openapi3.Schema, depth int) {
	required := map[string]bool{}
	for _, r := range v.Required {
		required[r] = true
	}
	names := make([]string, 0, len(v.Properties))
	for n := range v.Properties {
		names = append(names, n)
	}
	sort.Strings(names)
	indent := strings.Repeat("  ", depth)
	for _, name := range names {
		prop := v.Properties[name]
		fmt.Fprintf(b, "%s- `%s` (%s", indent, name, typeName(prop))
		if required[name] {
			b.WriteString(", required")
		}
		b.WriteString(")")
		if prop != nil && prop.Value != nil {
			if d := strings.TrimSpace(prop.Value.Description); d != "" {
				b.WriteString(": " + strings.ReplaceAll(d, "\n", " "))
			}
			if len(prop.Value.Enum) > 0 {
				fmt.Fprintf(b, " — one of %s", enumList(prop.Value.Enum))
			}
		}
		b.WriteString("\n")
		if depth+1 < maxSchemaDepth && prop != nil && prop.Value != nil {
			if len(prop.Value.Properties) > 0 {
				writeProperties(b, prop.Value, depth+1)
			} else if prop.Value.Items != nil && prop.Value.Items.Value != nil && len(prop.Value.Items.Value.Properties) > 0 {
				writeProperties(b, prop.Value.Items.Value, depth+1)
			}
		}
	}
}

// typeName gives a compact type label such as "string (date-time)",
// "array of integer" or "object".
func typeName(ref *openapi3.SchemaRef) string {
	if ref == nil || ref.Value == nil {
		return ""
	}
	v := ref.Value
	var t string
	if v.Type != nil {
		t = strings.Join(v.Type.Slice(), "|")
	}
	switch {
	case v.Items != nil:
		inner := typeName(v.Items)
		if inner == "" {
			inner = "any"
		}
		return "array of " + inner
	case t == "" && len(v.Properties) > 0:
		t = "object"
	case t == "" && len(v.OneOf) > 0:
		return "one of " + strings.Join(altNames(v.OneOf), " | ")
	case t == "" && len(v.AnyOf) > 0:
		return "any of " + strings.Join(altNames(v.AnyOf), " | ")
	case t == "" && len(v.AllOf) > 0:
		return "object"
	case t == "":
		return "any"
	}
	if v.Format != "" {
		t += " (" + v.Format + ")"
	}
	return t
}

func altNames(refs openapi3.SchemaRefs) []string {
	out := make([]string, 0, len(refs))
	for _, r := range refs {
		if r == nil {
			continue
		}
		if r.Ref != "" {
			out = append(out, r.Ref[strings.LastIndex(r.Ref, "/")+1:])
			continue
		}
		out = append(out, typeName(r))
	}
	return out
}

func enumList(vals []any) string {
	parts := make([]string, 0, len(vals))
	for _, v := range vals {
		parts = append(parts, fmt.Sprintf("`%v`", v))
	}
	return strings.Join(parts, ", ")
}

// writeExamples renders media-type examples as fenced code blocks: the single
// example if present, otherwise every named example.
func writeExamples(b *strings.Builder, ct string, media *openapi3.MediaType) {
	if media == nil {
		return
	}
	lang := "json"
	if !strings.Contains(ct, "json") {
		lang = ""
	}
	if media.Example != nil {
		b.WriteString("Example:\n\n")
		writeCode(b, lang, media.Example)
		return
	}
	if len(media.Examples) == 0 {
		if media.Schema != nil && media.Schema.Value != nil && media.Schema.Value.Example != nil {
			b.WriteString("Example:\n\n")
			writeCode(b, lang, media.Schema.Value.Example)
		} else if lang == "json" {
			if v := synthesize(media.Schema, 0); v != nil {
				b.WriteString("Example (generated from schema):\n\n")
				writeCode(b, lang, v)
			}
		}
		return
	}
	names := make([]string, 0, len(media.Examples))
	for n := range media.Examples {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		ex := media.Examples[n]
		if ex == nil || ex.Value == nil {
			continue
		}
		label := ex.Value.Summary
		if label == "" {
			label = n
		}
		fmt.Fprintf(b, "Example (%s):\n\n", label)
		writeCode(b, lang, ex.Value.Value)
	}
}

func writeCode(b *strings.Builder, lang string, v any) {
	var body string
	switch t := v.(type) {
	case string:
		body = t
	default:
		js, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			body = fmt.Sprint(v)
		} else {
			body = string(js)
		}
	}
	fmt.Fprintf(b, "```%s\n%s\n```\n\n", lang, strings.TrimSpace(body))
}

// synthesize builds an example value from a schema, using property-level
// examples where the spec provides them and type placeholders otherwise.
// It returns nil for schemas with no structure worth showing.
func synthesize(ref *openapi3.SchemaRef, depth int) any {
	if ref == nil || ref.Value == nil || depth > maxSchemaDepth {
		return nil
	}
	v := ref.Value
	if v.Example != nil {
		return v.Example
	}
	switch {
	case len(v.Properties) > 0:
		out := map[string]any{}
		for name, prop := range v.Properties {
			if val := synthesize(prop, depth+1); val != nil {
				out[name] = val
			}
		}
		if len(out) == 0 {
			return nil
		}
		return out
	case v.Items != nil:
		item := synthesize(v.Items, depth+1)
		if item == nil {
			return nil
		}
		return []any{item}
	case len(v.Enum) > 0:
		return v.Enum[0]
	}
	if v.Type == nil {
		return nil
	}
	switch {
	case v.Type.Is("string"):
		if v.Format != "" {
			return v.Format
		}
		return "string"
	case v.Type.Is("integer"):
		return 0
	case v.Type.Is("number"):
		return 0.0
	case v.Type.Is("boolean"):
		return true
	}
	return nil
}
