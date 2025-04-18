package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ErrorHandler wraps an http.Handler and provides consistent error handling
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer that captures the status code and body
		rw := &responseWriter{
			ResponseWriter: w,
			status:        http.StatusOK,
			body:          make([]byte, 0),
			headerWritten: false,
		}

		// Determine if we're in production based on hostname
		host := r.Host
		if colon := strings.Index(host, ":"); colon != -1 {
			host = host[:colon]
		}
		isProduction := host != "localhost" && host != "127.0.0.1"
		log.Printf("ErrorHandler: Host=%s (original=%s), isProduction=%v", host, r.Host, isProduction)

		// Recover from any panics
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				if !rw.headerWritten {
					if isProduction {
						http.Error(w, "Internal server error", http.StatusInternalServerError)
					} else {
						http.Error(w, fmt.Sprintf("Panic: %v", err), http.StatusInternalServerError)
					}
				}
			}
		}()

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Only log body content if it's text and not too long
		bodyStr := string(rw.body)
		if len(bodyStr) > 1000 {
			bodyStr = bodyStr[:1000] + "..."
		}

		// Check if the content is binary (contains non-printable characters)
		isBinary := false
		for _, b := range rw.body {
			if b < 32 && b != '\n' && b != '\r' && b != '\t' {
				isBinary = true
				break
			}
		}

		if !isBinary {
			log.Printf("ErrorHandler: Response status=%d, body=%q, containsError=%v",
				rw.status,
				bodyStr,
				containsError(rw.body))
		} else {
			log.Printf("ErrorHandler: Response status=%d, content-type=%s, body length=%d",
				rw.status,
				w.Header().Get("Content-Type"),
				len(rw.body))
		}

		// Check for error status codes or error messages in the body
		if rw.status >= http.StatusBadRequest || (rw.status == http.StatusOK && containsError(rw.body)) {
			log.Printf("Error %d for %s %s", rw.status, r.Method, r.URL.Path)
			if !rw.headerWritten {
				if isProduction {
					http.Error(w, "Internal server error", rw.status)
				} else {
					http.Error(w, fmt.Sprintf("Error %d: %s %s", rw.status, r.Method, r.URL.Path), rw.status)
				}
			}
		}
	})
}

// responseWriter is a custom response writer that captures the status code and body
type responseWriter struct {
	http.ResponseWriter
	status        int
	body          []byte
	headerWritten bool
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	if !rw.headerWritten {
		rw.status = code
		rw.headerWritten = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

// Write captures the body
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

// containsError checks if the response body contains an error message
func containsError(body []byte) bool {
	// Convert to string for easier processing
	content := string(body)

	// Skip checking if it's HTML content
	if strings.Contains(content, "<!DOCTYPE html>") ||
	   strings.Contains(content, "<html") ||
	   strings.Contains(content, "<head>") {
		return false
	}

	// Look for actual error messages
	errorPatterns := []string{
		"error:",
		"failed to",
		"panic:",
		"template:",
		"nil pointer",
		"invalid",
	}

	for _, pattern := range errorPatterns {
		if strings.Contains(strings.ToLower(content), pattern) {
			return true
		}
	}
	return false
}