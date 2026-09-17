// Package httpapi builds the Gin engine, wires middleware, and registers
// every HTTP route group.
package httpapi

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/shiprocket/apidocs/internal/auth"
	"github.com/shiprocket/apidocs/internal/config"
	"github.com/shiprocket/apidocs/internal/domain"
	"github.com/shiprocket/apidocs/internal/importsvc"
	"github.com/shiprocket/apidocs/internal/proxy"
)

// Pinger is the subset of the Mongo client the health check needs.
type Pinger interface {
	Ping(ctx context.Context) error
}

// Deps carries everything the router needs from main.
type Deps struct {
	Config      *config.Config
	Mongo       Pinger
	Logger      *slog.Logger
	Collections CollectionStore
	Operations  OperationStore
	Auth        *auth.Service
	Importer    *importsvc.Service
	Search      Searcher
	SiteName    string
	// UI is the built frontend bundle (see internal/ui). nil disables it.
	UI fs.FS
	// LoginPerMinute caps login attempts per client IP; 0 means the default of 5.
	LoginPerMinute int
	// ProxyPerMinute caps Try-it requests per client IP; 0 means the default of 60.
	ProxyPerMinute int
}

// NewRouter returns a configured *gin.Engine.
func NewRouter(d Deps) *gin.Engine {
	gin.SetMode(d.Config.GinMode)
	r := gin.New()
	r.Use(gin.Recovery(), requestLogger(d.Logger))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     d.Config.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/healthz", healthHandler(d.Mongo))

	public := r.Group("/api/v1")
	site := d.SiteName
	if site == "" {
		site = "API Docs"
	}
	app := newSPA(d.UI)
	var docsSPA gin.HandlerFunc
	if app.built() {
		docsSPA = app.serveIndex
	}
	docs := &docsHandlers{cols: d.Collections, search: d.Search, site: site, spa: docsSPA, log: d.Logger}
	docs.register(r, public)
	r.NoRoute(app.notFound)
	(&collectionHandlers{
		cols: d.Collections,
		ops:  d.Operations,
		log:  d.Logger,
		markdown: func(c *gin.Context) {
			slug := c.Param("slug")
			if !domain.ValidSlug(slug) {
				notFound(c)
				return
			}
			docs.serveMarkdown(c, slug, "", mimeMarkdown)
		},
	}).register(public)

	perMinute := d.LoginPerMinute
	if perMinute <= 0 {
		perMinute = 5
	}
	proxyPerMinute := d.ProxyPerMinute
	if proxyPerMinute <= 0 {
		proxyPerMinute = 60
	}
	tryIt := proxy.New(proxy.Config{AllowedHosts: d.Config.ProxyAllowedHosts, AllowPrivate: d.Config.ProxyAllowPrivate})
	public.Any("/proxy", newIPLimiter(proxyPerMinute).middleware(), tryIt.Gin())

	(&adminHandlers{
		cols:     d.Collections,
		ops:      d.Operations,
		importer: d.Importer,
		auth:     d.Auth,
		logins:   newIPLimiter(perMinute),
		log:      d.Logger,
	}).register(public)

	return r
}

func healthHandler(p Pinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := p.Ping(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "degraded", "mongo": "unreachable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "mongo": "ok"})
	}
}

// requestLogger emits one structured line per request without bodies or PII.
func requestLogger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("request",
			"method", c.Request.Method,
			"path", c.FullPath(),
			"status", c.Writer.Status(),
			"duration_ms", time.Since(start).Milliseconds(),
		)
	}
}
