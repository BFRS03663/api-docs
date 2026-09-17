package httpapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/shiprocket/apidocs/internal/auth"
	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importer"
	"github.com/shiprocket/apidocs/internal/importer/source"
	"github.com/shiprocket/apidocs/internal/importsvc"
)

// maxUploadBytes caps spec uploads; large real-world collections are ~2 MB.
const maxUploadBytes = 10 << 20

// CollectionStore is the full collection repository contract.
type CollectionStore interface {
	CollectionReader
	Upsert(ctx context.Context, c *domain.Collection) (bool, error)
	Delete(ctx context.Context, slug string) (string, error)
}

// OperationStore is the full operation repository contract.
type OperationStore interface {
	OperationReader
	Replace(ctx context.Context, collectionID string, ops []domain.Operation) error
	DeleteByCollection(ctx context.Context, collectionID string) error
}

type adminHandlers struct {
	cols     CollectionStore
	ops      OperationStore
	importer *importsvc.Service
	auth     *auth.Service
	logins   *ipLimiter
	log      logger
}

type logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

func (h *adminHandlers) register(public *gin.RouterGroup) {
	a := public.Group("/auth")
	a.POST("/login", h.logins.middleware(), h.login)
	a.POST("/logout", h.logout)
	a.GET("/me", h.auth.Middleware(), h.me)

	admin := public.Group("/admin", h.auth.Middleware())
	admin.POST("/collections", h.create)
	admin.PUT("/collections/:slug", h.reimport)
	admin.DELETE("/collections/:slug", h.remove)
}

type loginRequest struct {
	Username string `json:"username" binding:"required,max=128"`
	Password string `json:"password" binding:"required,max=1024"`
}

func (h *adminHandlers) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}
	token, exp, err := h.auth.Login(req.Username, req.Password)
	if err != nil {
		h.log.Info("admin login failed", "ip", c.ClientIP())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "expiresAt": exp.UTC().Format(time.RFC3339)})
}

// logout is stateless: tokens are short-lived and the client discards its
// copy. The endpoint exists so the UI has a uniform call to make.
func (h *adminHandlers) logout(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (h *adminHandlers) me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"username": c.GetString(auth.ContextKey)})
}

// create imports a new collection. 409 if the slug is already taken.
func (h *adminHandlers) create(c *gin.Context) {
	in, ok := h.readImport(c, "")
	if !ok {
		return
	}
	if _, err := h.cols.GetBySlug(c.Request.Context(), in.Slug); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "a collection with this slug already exists; use PUT to re-import"})
		return
	} else if !errors.Is(err, domain.ErrNotFound) {
		h.fail(c, err)
		return
	}
	h.runImport(c, in, http.StatusCreated)
}

// reimport replaces the spec of an existing collection.
func (h *adminHandlers) reimport(c *gin.Context) {
	slug := c.Param("slug")
	if !domain.ValidSlug(slug) {
		notFound(c)
		return
	}
	if _, err := h.cols.GetBySlug(c.Request.Context(), slug); err != nil {
		h.fail(c, err)
		return
	}
	in, ok := h.readImport(c, slug)
	if !ok {
		return
	}
	h.runImport(c, in, http.StatusOK)
}

func (h *adminHandlers) remove(c *gin.Context) {
	slug := c.Param("slug")
	if !domain.ValidSlug(slug) {
		notFound(c)
		return
	}
	id, err := h.cols.Delete(c.Request.Context(), slug)
	if err != nil {
		h.fail(c, err)
		return
	}
	if err := h.ops.DeleteByCollection(c.Request.Context(), id); err != nil {
		h.fail(c, err)
		return
	}
	h.log.Info("collection deleted", "slug", slug, "admin", c.GetString(auth.ContextKey))
	c.Status(http.StatusNoContent)
}

// readImport accepts either multipart/form-data (file, slug, name) or JSON
// ({url, slug, name}). fixedSlug overrides the body's slug for re-imports.
func (h *adminHandlers) readImport(c *gin.Context, fixedSlug string) (importsvc.Input, bool) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadBytes)
	var in importsvc.Input

	ct := c.ContentType()
	switch {
	case strings.HasPrefix(ct, "multipart/form-data"):
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			if tooLarge(err) {
				c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file exceeds 10 MB"})
				return in, false
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "multipart field 'file' is required"})
			return in, false
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			if tooLarge(err) {
				c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file exceeds 10 MB"})
				return in, false
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not read file"})
			return in, false
		}
		in.Data, in.Filename = data, header.Filename
		in.Slug, in.Name = c.PostForm("slug"), c.PostForm("name")
	case strings.HasPrefix(ct, "application/json"):
		var body struct {
			URL  string `json:"url" binding:"required,max=2048"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "json body with 'url' is required"})
			return in, false
		}
		in.URL, in.Slug, in.Name = body.URL, body.Slug, body.Name
	default:
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "send multipart/form-data with a file, or application/json with a url"})
		return in, false
	}

	if fixedSlug != "" {
		in.Slug = fixedSlug
	}
	in.Slug = strings.TrimSpace(in.Slug)
	in.Name = strings.TrimSpace(in.Name)
	if !domain.ValidSlug(in.Slug) {
		c.JSON(http.StatusBadRequest, gin.H{"error": importsvc.ErrInvalidSlug.Error()})
		return in, false
	}
	if len(in.Name) > 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is too long"})
		return in, false
	}
	return in, true
}

func (h *adminHandlers) runImport(c *gin.Context, in importsvc.Input, okStatus int) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 90*time.Second)
	defer cancel()
	col, created, err := h.importer.Import(ctx, in)
	if err != nil {
		switch {
		case errors.Is(err, importer.ErrUnknownFormat), errors.Is(err, importer.ErrInvalidSpec),
			errors.Is(err, importer.ErrUnsupportedFormat), errors.Is(err, source.ErrUnsupported):
			// Import errors describe the user's own upload; safe to return.
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			h.fail(c, err)
		}
		return
	}
	h.log.Info("collection imported", "slug", col.Slug, "created", created, "operations", col.OperationCount, "admin", c.GetString(auth.ContextKey))
	status := okStatus
	if created && status == http.StatusOK {
		status = http.StatusCreated
	}
	c.JSON(status, gin.H{"collection": col, "created": created})
}

func (h *adminHandlers) fail(c *gin.Context, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		notFound(c)
		return
	}
	h.log.Error("admin handler error", "path", c.FullPath(), "error", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
}

func tooLarge(err error) bool {
	var mbe *http.MaxBytesError
	return errors.As(err, &mbe)
}
