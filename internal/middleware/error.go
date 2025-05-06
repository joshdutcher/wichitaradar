package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// AppError is our main error type that includes HTTP context
type AppError struct {
	// Original error that caused this
	Err error
	// HTTP status code to return
	StatusCode int
	// Message to show in development
	DevMessage string
	// Message to show in production
	ProdMessage string
}

func (e *AppError) Error() string {
	return e.DevMessage
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// InternalError creates an internal server error
func InternalError(err error) *AppError {
	return &AppError{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
		DevMessage: fmt.Sprintf("Internal error: %v", err),
		ProdMessage: "Internal Server Error",
	}
}

// NotFoundError creates a not found error
func NotFoundError(err error, resource string) *AppError {
	return &AppError{
		Err:        err,
		StatusCode: http.StatusNotFound,
		DevMessage: fmt.Sprintf("%s not found: %v", resource, err),
		ProdMessage: "Not found",
	}
}

// BadRequestError creates a bad request error
func BadRequestError(err error, message string) *AppError {
	return &AppError{
		Err:        err,
		StatusCode: http.StatusBadRequest,
		DevMessage: fmt.Sprintf("Bad request: %s: %v", message, err),
		ProdMessage: message,
	}
}

// HandlerFunc is a function that handles an HTTP request and may return an error
type HandlerFunc func(http.ResponseWriter, *http.Request) error

// ErrorHandler wraps an http.Handler or HandlerFunc to provide error handling
func ErrorHandler(next interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer to capture the response
		cw := &customResponseWriter{ResponseWriter: w}

		// Handle panics
		defer func() {
			if err := recover(); err != nil {
				handleError(w, r, fmt.Errorf("panic: %v", err), "Internal server error", http.StatusInternalServerError)
			}
		}()

		// Call the next handler based on its type
		var handlerErr error
		switch h := next.(type) {
		case http.Handler:
			h.ServeHTTP(cw, r)
		case HandlerFunc:
			handlerErr = h(cw, r)
		case func(http.ResponseWriter, *http.Request) error:
			handlerErr = h(cw, r)
		case func(http.ResponseWriter, *http.Request):
			h(cw, r)
		default:
			handlerErr = fmt.Errorf("unsupported handler type: %T", next)
		}

		// Check for errors from the handler
		if handlerErr != nil {
			status := http.StatusInternalServerError
			message := "Internal server error"
			if appErr, ok := handlerErr.(*AppError); ok {
				status = appErr.StatusCode
				message = appErr.ProdMessage
			}
			handleError(w, r, handlerErr, message, status)
			return
		}

		// Check if the response contains an error
		if containsError(cw.body) {
			handleError(w, r, fmt.Errorf("error detected in response body"), "Internal server error", http.StatusInternalServerError)
			return
		}
	})
}

// customResponseWriter is a wrapper around http.ResponseWriter that captures the response
type customResponseWriter struct {
	http.ResponseWriter
	body []byte
}

// Write captures the response body
func (w *customResponseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}

// containsError checks if the response body contains an error message
func containsError(body []byte) bool {
	// Skip binary content
	if len(body) > 0 && body[0] == 0 {
		return false
	}

	// Convert to string for easier checking
	content := string(body)

	// Skip HTML content
	if strings.Contains(content, "<!DOCTYPE html>") || strings.Contains(content, "<html") {
		return false
	}

	// Check for error patterns
	errorPatterns := []string{
		"error",
		"failed",
		"not found",
		"internal server",
		"bad request",
		"unauthorized",
		"forbidden",
	}

	for _, pattern := range errorPatterns {
		if strings.Contains(strings.ToLower(content), pattern) {
			return true
		}
	}

	return false
}

// handleError writes an error response to the client
func handleError(w http.ResponseWriter, r *http.Request, err error, message string, status int) {
	// Determine if we're in production based on ENV variable
	isProduction := os.Getenv("ENV") == "production"

	// Only log errors in non-test environments
	if !strings.HasPrefix(r.Host, "127.0.0.1:") {
		log.Printf("Error: %v (status=%d, production=%v)", err, status, isProduction)
	}

	// Set the status code
	w.WriteHeader(status)

	// Write the error response
	if isProduction {
		http.Error(w, message, status)
	} else {
		http.Error(w, fmt.Sprintf("%s\n\nError: %v", message, err), status)
	}
}

// responseWriter is a custom response writer that captures the status code
type responseWriter struct {
	http.ResponseWriter
	status        int
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

// Write passes through to the underlying writer
func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}