// Package handlers contains HTTP request handlers for the application.
package handlers

import (
	"net/http"
)

// HandleHealth handles health check requests.
func HandleHealth(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
	return nil
}
