package httpapi

import (
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// spa serves the embedded single-page frontend: hashed assets with long
// cache lifetimes, every other unknown path gets index.html so client-side
// routes work on reload.
type spa struct {
	fsys  fs.FS
	index []byte
	files http.Handler
}

func newSPA(fsys fs.FS) *spa {
	if fsys == nil {
		return &spa{}
	}
	index, err := fs.ReadFile(fsys, "index.html")
	if err != nil {
		return &spa{}
	}
	return &spa{fsys: fsys, index: index, files: http.FileServerFS(fsys)}
}

func (s *spa) built() bool { return s.index != nil }

// serveIndex writes index.html, used for /docs/<slug> and client routes.
func (s *spa) serveIndex(c *gin.Context) {
	if !s.built() {
		c.JSON(http.StatusNotFound, gin.H{"error": "ui not built; run make build"})
		return
	}
	c.Header("Cache-Control", "no-cache")
	c.Data(http.StatusOK, "text/html; charset=utf-8", s.index)
}

// notFound is the router's fallback: JSON for API paths, static file when
// one exists in the bundle, otherwise the app shell.
func (s *spa) notFound(c *gin.Context) {
	p := c.Request.URL.Path
	if strings.HasPrefix(p, "/api/") || !s.built() || c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
		notFound(c)
		return
	}
	clean := strings.TrimPrefix(path.Clean(p), "/")
	if clean != "" && clean != "index.html" {
		if f, err := s.fsys.Open(clean); err == nil {
			f.Close()
			if strings.HasPrefix(clean, "assets/") {
				c.Header("Cache-Control", "public, max-age=31536000, immutable")
			}
			s.files.ServeHTTP(c.Writer, c.Request)
			return
		}
	}
	s.serveIndex(c)
}
