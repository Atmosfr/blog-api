package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Atmosfr/blog-api/internal/config"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func NewHealthHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(HealthResponse{
			Status:      "ok",
			Environment: cfg.Environment,
			Version:     cfg.Version,
		})
		if err != nil {
			slog.Error("Failed to encode health response", "error", err)
			return
		}
	}
	
}
