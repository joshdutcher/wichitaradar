package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"wichitaradar/internal/middleware"
	"wichitaradar/internal/testutils"
)

func TestHandleRainfall(t *testing.T) {
	// Initialize templates
	testutils.InitTemplates(t)

	// Create request
	req := httptest.NewRequest("GET", "/rainfall", nil)
	w := httptest.NewRecorder()

	// Execute handler
	middleware.ErrorHandler(HandleRainfall).ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check response body
	body := w.Body.String()
	if !bytes.Contains([]byte(body), []byte("Rainfall")) {
		t.Errorf("expected body to contain %q, got %q", "Rainfall", body)
	}
}