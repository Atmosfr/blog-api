package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Panic recovered", "error", err, "method", r.Method, "path", r.URL.Path, "ip", r.RemoteAddr, "stacktrace", string(debug.Stack()))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} ()

		next.ServeHTTP(w, r)
	})
}