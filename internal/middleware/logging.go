package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestID := GetRequestID(r.Context())

		rw := &ResponseWriter{ResponseWriter: w, statusCode: 0}

		next.ServeHTTP(rw, r)

		durationMS := time.Since(start).Milliseconds()
		slog.Info("Request completed", "request_id", requestID, "method", r.Method, "status", rw.statusCode, "path", r.URL.Path, "ip", r.RemoteAddr, "duration_ms", durationMS)
	})
}
