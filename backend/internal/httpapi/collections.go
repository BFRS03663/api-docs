package httpapi

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/shiprocket/apidocs/internal/domain"
)

// CollectionReader is the read side of the collection store.
type CollectionReader interface {
	List(ctx context.Context) ([]domain.Collection, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Collection, error)
	GetSpec(ctx context.Context, slug string) ([]byte, error)
}

// OperationReader lists the flattened operations of a collection.
type OperationReader interface {
	ListByCollection(ctx context.Context, collectionID string) ([]domain.Operation, error)
}

type collectionHandlers struct {
	cols     CollectionReader
	ops      OperationReader
	log      *slog.Logger
	markdown gin.HandlerFunc // serves the Markdown rendering for Accept: text/markdown
}

func (h *collectionHandlers) register(public *gin.RouterGroup) {
	public.GET("/collections", h.list)
	public.GET("/collections/:slug", h.get)
	public.GET("/collections/:slug/operations", h.listOperations)
}

func (h *collectionHandlers) list(c *gin.Context) {
	cols, err := h.cols.List(c.Request.Context())
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"collections": cols})
}

func (h *collectionHandlers) get(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept"), "text/markdown") && h.markdown != nil {
		h.markdown(c)
		return
	}
	col, ok := h.lookup(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, col)
}

func (h *collectionHandlers) listOperations(c *gin.Context) {
	col, ok := h.lookup(c)
	if !ok {
		return
	}
	ops, err := h.ops.ListByCollection(c.Request.Context(), col.ID)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"collection": col.Slug, "operations": ops})
}

// lookup validates the slug and loads the collection, writing the error
// response itself when it returns ok=false.
func (h *collectionHandlers) lookup(c *gin.Context) (*domain.Collection, bool) {
	slug := c.Param("slug")
	if !domain.ValidSlug(slug) {
		notFound(c)
		return nil, false
	}
	col, err := h.cols.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		h.fail(c, err)
		return nil, false
	}
	return col, true
}

func (h *collectionHandlers) fail(c *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		notFound(c)
		return
	}
	h.log.Error("handler error", "path", c.FullPath(), "error", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}
