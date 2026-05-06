package middleware

import (
	"log/slog"
	"net/http"
	"time"
)


func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		user_agent := r.Header.Get("User-Agent")

		if len(user_agent) > 200 {
			user_agent = user_agent[:200] + "..."
		}

		slog.Info("Incoming request", "method", r.Method, "path", r.URL.Path, "ip", r.RemoteAddr, "user_agent", user_agent)

		rw := &ResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)
		
		duration_ms := time.Since(start).Milliseconds()
		slog.Info("Request completed", "method", r.Method, "status", rw.statusCode, "path", r.URL.Path, "ip", r.RemoteAddr, "duration_ms", duration_ms, "user_agent", user_agent)
	})
}