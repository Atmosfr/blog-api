package main

import (
	"log/slog"
	"os"

	"github.com/Atmosfr/blog-api/internal/config"
	"github.com/Atmosfr/blog-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	slog.SetLogLoggerLevel(cfg.LogLevel)

	slog.Info("Config loaded", "port", cfg.Port, "environment", cfg.Environment, "log_level", cfg.LogLevel.String())

	app := server.NewApp(cfg)

	if err := app.Run(); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}
