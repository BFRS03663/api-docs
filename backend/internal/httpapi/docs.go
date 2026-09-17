package httpapi

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/export/markdown"
)

// Searcher runs full-text search over the operation index.
type Searcher interface {
	Search(ctx context.Context, q, collectionID string, limit int) ([]domain.SearchHit, error)
}

const (
	mimeMarkdown = "text/markdown; charset=utf-8"
	mimeText     = "text/plain; charset=utf-8"
)

// docsHandlers serves the text surface (llms.txt, Markdown, plain text),
// the raw spec files, and search.
type docsHandlers struct {
	cols   CollectionReader
	search Searcher
	site   string
	spa    gin.HandlerFunc // serves the UI for /docs/<slug> when set
	log    logger
}

func (h *docsHandlers) register(root *gin.Engine, public *gin.RouterGroup) {
	root.GET("/llms.txt", h.llmsIndex)
	root.GET("/llms-full.txt", h.llmsFull)
	root.GET("/docs/:slug", h.collection)
	root.GET("/docs/:slug/:file", h.file)
	public.GET("/search", h.searchHandler)
}

// baseURL reconstructs the public origin so text renderings carry absolute
// links, honouring reverse-proxy headers.
func baseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil || strings.EqualFold(c.GetHeader("X-Forwarded-Proto"), "https") {
		scheme = "https"
	}
	host := c.GetHeader("X-Forwarded-Host")
	if host == "" {
		host = c.Request.Host
	}
	return scheme + "://" + host
}

func (h *docsHandlers) llmsIndex(c *gin.Context) {
	cols, err := h.cols.List(c.Request.Context())
	if err != nil {
		h.fail(c, err)
		return
	}
	c.Data(http.StatusOK, mimeText, []byte(markdown.Index(h.site, baseURL(c), cols)))
}

func (h *docsHandlers) llmsFull(c *gin.Context) {
	cols, err := h.cols.List(c.Request.Context())
	if err != nil {
		h.fail(c, err)
		return
	}
	var b strings.Builder
	b.WriteString(markdown.Index(h.site, baseURL(c), cols))
	for _, col := range cols {
		md, err := h.render(c, col.Slug, "")
		if err != nil {
			h.fail(c, err)
			return
		}
		b.WriteString("\n\n---\n\n")
		b.WriteString(md)
	}
	c.Data(http.StatusOK, mimeText, []byte(b.String()))
}

// collection handles /docs/<slug>.md, /docs/<slug>.txt, and the bare
// /docs/<slug> (Markdown when the client asks for it, otherwise the UI).
func (h *docsHandlers) collection(c *gin.Context) {
	name := c.Param("slug")
	slug, ext := splitExt(name)
	if !domain.ValidSlug(slug) {
		notFound(c)
		return
	}
	wantsMarkdown := ext == ".md" || (ext == "" && strings.Contains(c.GetHeader("Accept"), "text/markdown"))
	switch {
	case wantsMarkdown:
		h.serveMarkdown(c, slug, "", mimeMarkdown)
	case ext == ".txt":
		h.serveMarkdown(c, slug, "", mimeText)
	case ext == "" && h.spa != nil:
		h.spa(c)
	default:
		notFound(c)
	}
}

// file handles /docs/<slug>/openapi.json, /docs/<slug>/original and
// /docs/<slug>/<operationId>.md|.txt.
func (h *docsHandlers) file(c *gin.Context) {
	slug := c.Param("slug")
	if !domain.ValidSlug(slug) {
		notFound(c)
		return
	}
	switch file := c.Param("file"); file {
	case "openapi.json":
		spec, err := h.cols.GetSpec(c.Request.Context(), slug)
		if err != nil {
			h.fail(c, err)
			return
		}
		c.Header("Cache-Control", "public, max-age=60")
		c.Data(http.StatusOK, "application/json; charset=utf-8", spec)
	default:
		opID, ext := splitExt(file)
		if opID == "" || (ext != ".md" && ext != ".txt") {
			notFound(c)
			return
		}
		mime := mimeMarkdown
		if ext == ".txt" {
			mime = mimeText
		}
		h.serveMarkdown(c, slug, opID, mime)
	}
}

func (h *docsHandlers) serveMarkdown(c *gin.Context, slug, opID, mime string) {
	md, err := h.render(c, slug, opID)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.Header("Cache-Control", "public, max-age=60")
	c.Data(http.StatusOK, mime, []byte(md))
}

// render produces the Markdown for a collection, or for one operation when
// opID is set.
func (h *docsHandlers) render(c *gin.Context, slug, opID string) (string, error) {
	ctx := c.Request.Context()
	col, err := h.cols.GetBySlug(ctx, slug)
	if err != nil {
		return "", err
	}
	spec, err := h.cols.GetSpec(ctx, slug)
	if err != nil {
		return "", err
	}
	doc, err := markdown.Load(spec)
	if err != nil {
		return "", err
	}
	if opID == "" {
		return markdown.Collection(*col, doc, baseURL(c)), nil
	}
	md, ok := markdown.Operation(*col, doc, opID, baseURL(c))
	if !ok {
		return "", domain.ErrNotFound
	}
	return md, nil
}

func (h *docsHandlers) searchHandler(c *gin.Context) {
	q := strings.TrimSpace(c.Query("q"))
	if q == "" || len(q) > 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "q is required (1-200 characters)"})
		return
	}
	limit := 25
	if raw := c.Query("limit"); raw != "" {
		n, err := strconv.Atoi(raw)
		if err != nil || n < 1 || n > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be between 1 and 100"})
			return
		}
		limit = n
	}
	ctx := c.Request.Context()
	cols, err := h.cols.List(ctx)
	if err != nil {
		h.fail(c, err)
		return
	}
	byID := map[string]domain.Collection{}
	for _, col := range cols {
		byID[col.ID] = col
	}

	collectionID := ""
	if slug := c.Query("collection"); slug != "" {
		if !domain.ValidSlug(slug) {
			notFound(c)
			return
		}
		col, err := h.cols.GetBySlug(ctx, slug)
		if err != nil {
			h.fail(c, err)
			return
		}
		collectionID = col.ID
	}

	hits, err := h.search.Search(ctx, q, collectionID, limit)
	if err != nil {
		h.fail(c, err)
		return
	}
	for i := range hits {
		if col, ok := byID[hits[i].CollectionID]; ok {
			hits[i].CollectionSlug, hits[i].CollectionName = col.Slug, col.Name
		}
	}
	c.JSON(http.StatusOK, gin.H{"query": q, "hits": hits})
}

func (h *docsHandlers) fail(c *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		notFound(c)
		return
	}
	h.log.Error("docs handler error", "path", c.FullPath(), "error", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
}

// splitExt separates "petstore.md" into ("petstore", ".md"); names without a
// recognised extension come back unchanged.
func splitExt(name string) (string, string) {
	for _, ext := range []string{".md", ".txt"} {
		if strings.HasSuffix(name, ext) {
			return strings.TrimSuffix(name, ext), ext
		}
	}
	return name, ""
}
