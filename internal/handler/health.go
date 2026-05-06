package handler

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(HealthResponse{
		Status:      "ok",
		Environment: "...",
		Version:     "1.0.0",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
