package middleware

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}