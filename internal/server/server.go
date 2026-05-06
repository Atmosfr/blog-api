package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Atmosfr/blog-api/internal/config"
	"github.com/Atmosfr/blog-api/internal/handler"
)

type App struct {
	srv    *http.Server
	Config *config.Config
}

func (app *App) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	slog.Info("Starting server...")
	go func() {
		if err := app.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server error", "error", err)
		}
	}()

	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.srv.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error", "error", err)
		return err
	}
	slog.Info("Server gracefully stopped")

	return nil
}

func NewApp(config *config.Config) *App {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthHandler)

	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	return &App{
		srv:    server,
		Config: config,
	}
}
