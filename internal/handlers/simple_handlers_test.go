package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"wichitaradar/internal/middleware"
	"wichitaradar/internal/testutils"
)

func TestHandleSimplePage(t *testing.T) {
	tests := []struct {
		name           string
		templateName   string
		path           string
		setupTemplates func(*testing.T)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:         "about page success",
			templateName: "about",
			path:         "/about",
			setupTemplates: func(t *testing.T) {
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "About",
		},
		{
			name:         "disclaimer page success",
			templateName: "disclaimer",
			path:         "/disclaimer",
			setupTemplates: func(t *testing.T) {
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Disclaimer",
		},
		{
			name:         "donate page success",
			templateName: "donate",
			path:         "/donate",
			setupTemplates: func(t *testing.T) {
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Donate",
		},
		{
			name:         "resources page success",
			templateName: "resources",
			path:         "/resources",
			setupTemplates: func(t *testing.T) {
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Resources",
		},
		{
			name:         "watches page success",
			templateName: "watches",
			path:         "/watches",
			setupTemplates: func(t *testing.T) {
				testutils.InitTemplates(t)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Watches",
		},
		{
			name:         "template error",
			templateName: "nonexistent",
			path:         "/nonexistent",
			setupTemplates: func(t *testing.T) {
				// Don't initialize templates to force an error
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
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			// Execute handler
			middleware.ErrorHandler(HandleSimplePage(tt.templateName)).ServeHTTP(w, req)

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