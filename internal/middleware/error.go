package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler wraps an http.Handler and provides consistent error handling
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer that captures the status code
		rw := &responseWriter{
			ResponseWriter: w,
			status:        http.StatusOK,
		}

		// Determine if we're in production based on hostname
		isProduction := r.Host != "localhost" && r.Host != "127.0.0.1"

		// Recover from any panics
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				if isProduction {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				} else {
					http.Error(w, fmt.Sprintf("Panic: %v", err), http.StatusInternalServerError)
				}
			}
		}()

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Check for error status codes
		if rw.status >= http.StatusBadRequest {
			log.Printf("Error %d for %s %s", rw.status, r.Method, r.URL.Path)
			if isProduction {
				http.Error(w, "Internal server error", rw.status)
			} else {
				http.Error(w, fmt.Sprintf("Error %d: %s %s", rw.status, r.Method, r.URL.Path), rw.status)
			}
		}
	})
}

// responseWriter is a custom response writer that captures the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}