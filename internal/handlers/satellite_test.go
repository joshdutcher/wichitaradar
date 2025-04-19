package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"wichitaradar/internal/middleware"
	"wichitaradar/internal/testutils"
)

func TestHandleSatellite(t *testing.T) {
	// Initialize templates
	testutils.InitTemplates(t)

	// Create request
	req := httptest.NewRequest("GET", "/satellite", nil)
	w := httptest.NewRecorder()

	// Execute handler
	middleware.ErrorHandler(HandleSatellite).ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check response body
	body := w.Body.String()
	if !bytes.Contains([]byte(body), []byte("Satellite")) {
		t.Errorf("expected body to contain %q, got %q", "Satellite", body)
	}
}