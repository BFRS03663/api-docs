package markdown

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
)

// maxIntroChars caps the collection introduction quoted in llms.txt so the
// index stays an orientation, not a copy of the full rendering.
const maxIntroChars = 700

// maxGuideHeadings caps how many guide chapters are named per collection.
const maxGuideHeadings = 10

// IndexEntry pairs a collection with its parsed spec. Doc may be nil when the
// spec could not be loaded; the entry then renders as a bare link.
type IndexEntry struct {
	Collection domain.Collection
	Doc        *openapi3.T
}

// Index renders llms.txt (llmstxt.org convention): a site orientation, a
// quick list of collections, then one section per collection with its base
// URLs, authentication, introduction and every endpoint grouped by tag. Each
// endpoint links to its own Markdown page so a reader can fetch just the
// detail it needs.
func Index(siteName, baseURL string, entries []IndexEntry) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", siteName)
	fmt.Fprintf(&b, "> API reference for %d collection(s). This file lists every documented endpoint with its base URL and authentication. Each endpoint links to a Markdown page with its parameters, request body, responses and examples; each collection links to its complete documentation and its OpenAPI document.\n\n", len(entries))

	writeHowToRead(&b, baseURL)

	b.WriteString("## Collections\n\n")
	for _, e := range entries {
		writeCollectionLine(&b, baseURL, e.Collection)
	}
	b.WriteString("\n")

	for _, e := range entries {
		writeIndexCollection(&b, baseURL, e)
	}

	b.WriteString("## Optional\n\n")
	fmt.Fprintf(&b, "- [Everything in one file](%s/llms-full.txt): all collections concatenated, with full request and response details\n", baseURL)
	fmt.Fprintf(&b, "- [Search](%s/api/v1/search?q=): full-text search over endpoints; add `&collection=<slug>` to narrow to one API\n", baseURL)
	return b.String()
}

func writeHowToRead(b *strings.Builder, baseURL string) {
	b.WriteString("## How to read this file\n\n")
	fmt.Fprintf(b, "- `%s/docs/<collection>.md`: complete documentation for one collection (guides, every endpoint, schemas, examples)\n", baseURL)
	fmt.Fprintf(b, "- `%s/docs/<collection>/<operationId>.md`: one endpoint on its own\n", baseURL)
	fmt.Fprintf(b, "- `%s/docs/<collection>/openapi.json`: the machine-readable OpenAPI 3 document\n", baseURL)
	b.WriteString("- Endpoint paths below are relative to the collection's base URL. Replace `{placeholders}` with real values.\n\n")
}

func writeCollectionLine(b *strings.Builder, baseURL string, col domain.Collection) {
	line := firstLine(col.Description)
	if line != "" {
		line = ": " + line
	}
	fmt.Fprintf(b, "- [%s](%s/docs/%s.md)%s (%d endpoints)\n", col.Name, baseURL, col.Slug, line, col.OperationCount)
}

// writeIndexCollection renders the detailed section for one collection.
func writeIndexCollection(b *strings.Builder, baseURL string, e IndexEntry) {
	col := e.Collection
	fmt.Fprintf(b, "## %s\n\n", col.Name)
	if intro := firstParagraph(col.Description); intro != "" {
		b.WriteString(intro)
		b.WriteString("\n\n")
	}
	if col.Version != "" {
		fmt.Fprintf(b, "- Version: %s\n", col.Version)
	}
	if len(col.Servers) > 0 {
		fmt.Fprintf(b, "- Base URL: %s\n", strings.Join(col.Servers, ", "))
	}
	if auth := securitySummary(e.Doc); auth != "" {
		fmt.Fprintf(b, "- Authentication: %s\n", auth)
	}
	if guides := guideHeadings(col.Description); len(guides) > 0 {
		fmt.Fprintf(b, "- Guides in the full documentation: %s\n", strings.Join(guides, "; "))
	}
	fmt.Fprintf(b, "- Endpoints: %d\n", col.OperationCount)
	fmt.Fprintf(b, "- [Full documentation](%s/docs/%s.md) | [OpenAPI document](%s/docs/%s/openapi.json)\n\n", baseURL, col.Slug, baseURL, col.Slug)

	if e.Doc == nil {
		return
	}
	groups, order := groupByTag(e.Doc, importer.Flatten(e.Doc))
	for _, tag := range order {
		g := groups[tag]
		fmt.Fprintf(b, "### %s\n\n", g.display)
		if d := firstParagraph(g.description); d != "" {
			b.WriteString(d)
			b.WriteString("\n\n")
		}
		for _, op := range g.ops {
			writeEndpointLine(b, baseURL, col.Slug, op)
		}
		b.WriteString("\n")
	}
}

func writeEndpointLine(b *strings.Builder, baseURL, slug string, op domain.Operation) {
	fmt.Fprintf(b, "- [%s %s](%s/docs/%s/%s.md)", op.Method, op.Path, baseURL, slug, op.OperationID)
	if s := firstLine(op.Summary); s != "" {
		fmt.Fprintf(b, ": %s", s)
	}
	if op.Deprecated {
		b.WriteString(" (deprecated)")
	}
	b.WriteString("\n")
}

// securitySummary describes the spec's security schemes in one line, e.g.
// "HTTP bearer token in the Authorization header (bearerAuth)".
func securitySummary(doc *openapi3.T) string {
	if doc == nil || doc.Components == nil || len(doc.Components.SecuritySchemes) == 0 {
		return ""
	}
	names := make([]string, 0, len(doc.Components.SecuritySchemes))
	for name := range doc.Components.SecuritySchemes {
		names = append(names, name)
	}
	sort.Strings(names)
	var parts []string
	for _, name := range names {
		ref := doc.Components.SecuritySchemes[name]
		if ref == nil || ref.Value == nil {
			continue
		}
		parts = append(parts, describeScheme(name, ref.Value))
	}
	return strings.Join(parts, "; ")
}

func describeScheme(name string, s *openapi3.SecurityScheme) string {
	var desc string
	switch s.Type {
	case "http":
		if strings.EqualFold(s.Scheme, "bearer") {
			desc = "HTTP bearer token in the Authorization header"
		} else {
			desc = "HTTP " + s.Scheme + " authentication"
		}
	case "apiKey":
		desc = fmt.Sprintf("API key `%s` in the %s", s.Name, s.In)
	case "oauth2":
		desc = "OAuth 2.0"
	case "openIdConnect":
		desc = "OpenID Connect"
	default:
		desc = s.Type
	}
	return fmt.Sprintf("%s (%s)", desc, name)
}

// guideHeadings lists the top-level chapter titles of a collection's
// description (Markdown "#" and "##" headings), so a reader knows which
// guides the full documentation contains before fetching it.
func guideHeadings(description string) []string {
	var out []string
	seen := map[string]bool{}
	for line := range strings.SplitSeq(description, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "# ") && !strings.HasPrefix(line, "## ") {
			continue
		}
		title := strings.TrimSpace(strings.TrimLeft(line, "# "))
		title = strings.Trim(title, "*:")
		if title == "" || seen[title] {
			continue
		}
		seen[title] = true
		out = append(out, title)
		if len(out) == maxGuideHeadings {
			break
		}
	}
	return out
}

// firstParagraph returns the first non-heading paragraph of Markdown text,
// collapsed onto one line and capped at maxIntroChars.
func firstParagraph(s string) string {
	for para := range strings.SplitSeq(strings.ReplaceAll(strings.TrimSpace(s), "\r\n", "\n"), "\n\n") {
		var lines []string
		for line := range strings.SplitSeq(para, "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			lines = append(lines, strings.TrimLeft(line, "> "))
		}
		if len(lines) == 0 {
			continue
		}
		text := strings.Join(lines, " ")
		if len(text) > maxIntroChars {
			text = text[:maxIntroChars-3] + "..."
		}
		return text
	}
	return ""
}
