// Command server runs the API documentation portal backend.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shiprocket/apidocs/internal/auth"
	"github.com/shiprocket/apidocs/internal/config"
	"github.com/shiprocket/apidocs/internal/httpapi"
	"github.com/shiprocket/apidocs/internal/importer/source"
	"github.com/shiprocket/apidocs/internal/importsvc"
	"github.com/shiprocket/apidocs/internal/repo/mongo"
	"github.com/shiprocket/apidocs/internal/ui"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	db, err := mongo.Connect(connectCtx, cfg.MongoURI, cfg.MongoDB)
	cancel()
	if err != nil {
		logger.Error("mongo unavailable", "error", err)
		os.Exit(1)
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := db.Close(closeCtx); err != nil {
			logger.Warn("mongo disconnect", "error", err)
		}
	}()

	cols := mongo.NewCollectionRepo(db.DB)
	ops := mongo.NewOperationRepo(db.DB)
	indexCtx, cancelIndex := context.WithTimeout(ctx, 30*time.Second)
	defer cancelIndex()
	if err := cols.EnsureIndexes(indexCtx); err != nil {
		logger.Error("collection indexes", "error", err)
		os.Exit(1)
	}
	if err := ops.EnsureIndexes(indexCtx); err != nil {
		logger.Error("operation indexes", "error", err)
		os.Exit(1)
	}

	router := httpapi.NewRouter(httpapi.Deps{
		Config:      cfg,
		Mongo:       db,
		Logger:      logger,
		Collections: cols,
		Operations:  ops,
		Auth:        auth.New(cfg.AdminUsername, cfg.AdminPasswordHash, cfg.JWTSecret, 12*time.Hour),
		Importer:    importsvc.New(cols, ops, source.New(cfg.ProxyAllowPrivate)),
		Search:      ops,
		SiteName:    cfg.SiteName,
		UI:          ui.FS(),
	})
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	go func() {
		logger.Info("server listening", "port", cfg.Port, "mode", cfg.GinMode)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server failed", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down")
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancelShutdown()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("graceful shutdown failed", "error", err)
	}
}
