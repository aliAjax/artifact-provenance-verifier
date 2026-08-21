package main

import (
	"context"
	"github.com/example/artifact-provenance-verifier/internal/api"
	"github.com/example/artifact-provenance-verifier/internal/platform"
	"github.com/example/artifact-provenance-verifier/internal/storage/infrastructure"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := platform.LoadConfig()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	repo := infrastructure.NewMemory()
	srv := &http.Server{Addr: cfg.HTTPAddr, Handler: api.New(cfg, repo, log).Handler()}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		log.Info("server started", "addr", cfg.HTTPAddr)
		if e := srv.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			log.Error("server failed", "error", e)
		}
	}()
	<-ctx.Done()
	log.Info("shutdown requested")
	if e := api.Shutdown(context.Background(), srv); e != nil {
		log.Error("shutdown failed", "error", e)
	}
}
