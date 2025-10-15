package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"wichitaradar/internal/middleware"
	"wichitaradar/internal/testutils"
	"wichitaradar/pkg/templates"
)

func TestHandleHome(t *testing.T) {
	tests := []struct {
		name           string
		setupTemplates func(*testing.T)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "successful render",
			setupTemplates: func(t *testing.T) {
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Wichita Radar",
		},
		{
			name: "template error",
			setupTemplates: func(t *testing.T) {
				// Reset templates to force an error
				templates.Reset()
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.setupTemplates != nil {
				tt.setupTemplates(t)
			}

			// Create request
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()

			// Execute handler
			middleware.ErrorHandler(HandleHome).ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check response body
			body := w.Body.String()
			if !bytes.Contains([]byte(body), []byte(tt.expectedBody)) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedBody, body)
			}
		})
	}
}
