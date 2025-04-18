package middleware

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// ErrorHandler wraps an http.Handler and provides consistent error handling
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer that captures the status code and body
		rw := &responseWriter{
			ResponseWriter: w,
			status:        http.StatusOK,
			body:          make([]byte, 0),
		}

		// Determine if we're in production based on hostname
		isProduction := r.Host != "localhost" && r.Host != "127.0.0.1"
		log.Printf("ErrorHandler: Host=%s, isProduction=%v", r.Host, isProduction)

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

		// Log the response state
		log.Printf("ErrorHandler: Response status=%d, body=%q, containsError=%v",
			rw.status,
			string(rw.body),
			containsError(rw.body))

		// Check for error status codes or error messages in the body
		if rw.status >= http.StatusBadRequest || (rw.status == http.StatusOK && containsError(rw.body)) {
			log.Printf("Error %d for %s %s", rw.status, r.Method, r.URL.Path)
			if isProduction {
				http.Error(w, "Internal server error", rw.status)
			} else {
				http.Error(w, fmt.Sprintf("Error %d: %s %s", rw.status, r.Method, r.URL.Path), rw.status)
			}
		}
	})
}

// responseWriter is a custom response writer that captures the status code and body
type responseWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures the body
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

// containsError checks if the response body contains an error message
func containsError(body []byte) bool {
	errorKeywords := []string{"error", "failed", "panic", "template", "nil"}
	for _, keyword := range errorKeywords {
		if bytes.Contains(body, []byte(keyword)) {
			return true
		}
	}
	return false
}