package handlers

import (
	"net/http"
)

// HandleHealth handles health check requests
func HandleHealth(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	return nil
}