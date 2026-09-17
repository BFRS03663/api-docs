package postman

import (
	"regexp"
	"strings"

	htmltomd "github.com/JohannesKaufmann/html-to-markdown/v2"
	mdconv "github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/table"
)

// htmlMarker matches the tags Postman's published-docs export uses when it
// pre-renders descriptions to HTML. Plain Markdown never contains them.
var htmlMarker = regexp.MustCompile(`(?i)<(html|body|p|div|table|h[1-6]|ul|ol|pre|br)\b`)

var mdConverter = mdconv.NewConverter(
	mdconv.WithPlugins(
		base.NewBasePlugin(),
		commonmark.NewCommonmarkPlugin(),
		table.NewTablePlugin(),
	),
)

// toMarkdown returns the description as CommonMark. HTML (as produced by
// Postman's documenter export) is converted; Markdown or plain text is
// returned trimmed and unchanged.
func toMarkdown(s string) string {
	s = strings.TrimSpace(s)
	if s == "" || !htmlMarker.MatchString(s) {
		return s
	}
	md, err := mdConverter.ConvertString(s)
	if err != nil {
		// Fall back to a tag-stripped version rather than losing the text.
		md, _ = htmltomd.ConvertString(s)
	}
	return strings.TrimSpace(md)
}
