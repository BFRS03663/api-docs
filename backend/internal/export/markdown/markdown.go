// Package markdown renders stored OpenAPI documents as Markdown and plain
// text so LLMs, crawlers and humans without a browser can read the same
// documentation the UI shows.
package markdown

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
)

// Load parses a stored canonical spec without validation (it was validated
// on import) and with local $refs resolved.
func Load(spec []byte) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = false
	return loader.LoadFromData(spec)
}

// Collection renders a whole collection: header, description, then every
// operation grouped under its tags in the spec's tag order. baseURL is used
// for absolute links and may be empty.
func Collection(col domain.Collection, doc *openapi3.T, baseURL string) string {
	var b strings.Builder
	writeHeader(&b, col, doc, baseURL)

	ops := importer.Flatten(doc)
	groups, order := groupByTag(doc, ops)
	for _, tag := range order {
		g := groups[tag]
		fmt.Fprintf(&b, "## %s\n\n", g.display)
		if g.description != "" {
			b.WriteString(strings.TrimSpace(g.description))
			b.WriteString("\n\n")
		}
		for _, op := range g.ops {
			writeOperation(&b, doc, op, 3)
		}
	}
	return strings.TrimSpace(b.String()) + "\n"
}

// Operation renders one operation with just enough collection context to
// stand alone. ok is false when operationID is unknown.
func Operation(col domain.Collection, doc *openapi3.T, operationID string, baseURL string) (string, bool) {
	for _, op := range importer.Flatten(doc) {
		if op.OperationID != operationID {
			continue
		}
		var b strings.Builder
		fmt.Fprintf(&b, "# %s\n\n", col.Name)
		writeServers(&b, doc)
		if baseURL != "" {
			fmt.Fprintf(&b, "Full documentation: %s/docs/%s.md\n\n", baseURL, col.Slug)
		}
		writeOperation(&b, doc, op, 2)
		return strings.TrimSpace(b.String()) + "\n", true
	}
	return "", false
}

type tagGroup struct {
	display     string
	description string
	ops         []domain.Operation
}

// groupByTag buckets operations by tag, honouring the spec's tag order and
// x-displayName, then appending undeclared tags and finally untagged ops.
func groupByTag(doc *openapi3.T, ops []domain.Operation) (map[string]*tagGroup, []string) {
	groups := map[string]*tagGroup{}
	var order []string
	for _, t := range doc.Tags {
		if t == nil || t.Name == "" {
			continue
		}
		display := t.Name
		if dn, ok := t.Extensions["x-displayName"].(string); ok && dn != "" {
			display = dn
		}
		groups[t.Name] = &tagGroup{display: display, description: t.Description}
		order = append(order, t.Name)
	}
	const untagged = "\x00untagged"
	for _, op := range ops {
		tags := op.Tags
		if len(tags) == 0 {
			tags = []string{untagged}
		}
		for _, tag := range tags {
			g, ok := groups[tag]
			if !ok {
				display := tag
				if tag == untagged {
					display = "Other endpoints"
				}
				g = &tagGroup{display: display}
				groups[tag] = g
				order = append(order, tag)
			}
			g.ops = append(g.ops, op)
		}
	}
	// Drop declared tags that ended up empty.
	kept := order[:0]
	for _, tag := range order {
		if len(groups[tag].ops) > 0 {
			kept = append(kept, tag)
		}
	}
	return groups, kept
}

func writeHeader(b *strings.Builder, col domain.Collection, doc *openapi3.T, baseURL string) {
	fmt.Fprintf(b, "# %s\n\n", col.Name)
	if col.Version != "" {
		fmt.Fprintf(b, "Version: %s\n\n", col.Version)
	}
	writeServers(b, doc)
	if baseURL != "" {
		fmt.Fprintf(b, "OpenAPI document: %s/docs/%s/openapi.json\n\n", baseURL, col.Slug)
	}
	if doc.Info != nil && strings.TrimSpace(doc.Info.Description) != "" {
		b.WriteString(strings.TrimSpace(doc.Info.Description))
		b.WriteString("\n\n")
	}
}

func writeServers(b *strings.Builder, doc *openapi3.T) {
	var urls []string
	for _, s := range doc.Servers {
		if s != nil && s.URL != "" {
			urls = append(urls, s.URL)
		}
	}
	if len(urls) > 0 {
		fmt.Fprintf(b, "Base URL: %s\n\n", strings.Join(urls, ", "))
	}
}

// writeOperation renders one operation; level is the heading depth to use.
func writeOperation(b *strings.Builder, doc *openapi3.T, op domain.Operation, level int) {
	h := strings.Repeat("#", level)
	sub := strings.Repeat("#", level+1)
	item := doc.Paths.Value(op.Path)
	if item == nil {
		return
	}
	o := item.Operations()[op.Method]
	if o == nil {
		return
	}

	fmt.Fprintf(b, "%s %s %s\n\n", h, op.Method, op.Path)
	if op.Summary != "" {
		fmt.Fprintf(b, "**%s**\n\n", strings.TrimSpace(op.Summary))
	}
	if op.Deprecated {
		b.WriteString("_Deprecated._\n\n")
	}
	if d := strings.TrimSpace(op.Description); d != "" {
		b.WriteString(d)
		b.WriteString("\n\n")
	}
	fmt.Fprintf(b, "Operation ID: `%s`\n\n", op.OperationID)

	if params := allParameters(item, o); len(params) > 0 {
		fmt.Fprintf(b, "%s Parameters\n\n", sub)
		b.WriteString("| Name | In | Required | Type | Description |\n|---|---|---|---|---|\n")
		for _, p := range params {
			req := "no"
			if p.Required {
				req = "yes"
			}
			fmt.Fprintf(b, "| `%s` | %s | %s | %s | %s |\n", p.Name, p.In, req, typeName(p.Schema), cell(p.Description))
		}
		b.WriteString("\n")
	}

	if o.RequestBody != nil && o.RequestBody.Value != nil {
		for _, ct := range sortedKeys(o.RequestBody.Value.Content) {
			media := o.RequestBody.Value.Content[ct]
			fmt.Fprintf(b, "%s Request body (%s)\n\n", sub, ct)
			if d := strings.TrimSpace(o.RequestBody.Value.Description); d != "" {
				b.WriteString(d)
				b.WriteString("\n\n")
			}
			writeSchema(b, media.Schema)
			writeExamples(b, ct, media)
		}
	}

	if o.Responses != nil && o.Responses.Len() > 0 {
		fmt.Fprintf(b, "%s Responses\n\n", sub)
		codes := o.Responses.Keys()
		sort.Strings(codes)
		for _, code := range codes {
			r := o.Responses.Value(code)
			if r == nil || r.Value == nil {
				continue
			}
			desc := ""
			if r.Value.Description != nil {
				desc = strings.TrimSpace(*r.Value.Description)
			}
			fmt.Fprintf(b, "**%s**", code)
			if desc != "" {
				fmt.Fprintf(b, " — %s", desc)
			}
			b.WriteString("\n\n")
			for _, ct := range sortedKeys(r.Value.Content) {
				media := r.Value.Content[ct]
				writeSchema(b, media.Schema)
				writeExamples(b, ct, media)
			}
		}
	}
}

func allParameters(item *openapi3.PathItem, o *openapi3.Operation) []*openapi3.Parameter {
	var out []*openapi3.Parameter
	seen := map[string]bool{}
	add := func(refs openapi3.Parameters) {
		for _, ref := range refs {
			if ref == nil || ref.Value == nil {
				continue
			}
			key := ref.Value.In + ":" + ref.Value.Name
			if seen[key] {
				continue
			}
			seen[key] = true
			out = append(out, ref.Value)
		}
	}
	add(o.Parameters)
	add(item.Parameters)
	return out
}

func sortedKeys(c openapi3.Content) []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// cell makes free text safe inside a Markdown table cell.
func cell(s string) string {
	s = strings.ReplaceAll(strings.TrimSpace(s), "\n", " ")
	return strings.ReplaceAll(s, "|", "\\|")
}

func firstLine(s string) string {
	for line := range strings.SplitSeq(strings.TrimSpace(s), "\n") {
		line = strings.TrimSpace(strings.TrimLeft(line, "#> "))
		if line != "" {
			if len(line) > 160 {
				line = line[:157] + "..."
			}
			return line
		}
	}
	return ""
}
