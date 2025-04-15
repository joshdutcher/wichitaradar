package testutils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"wichitaradar/pkg/templates"
)

// InitTemplates initializes the templates for testing
func InitTemplates(t *testing.T) {
	// Get the project root directory by going up from the current test file
	projectRoot, err := filepath.Abs("../../")
	if err != nil {
		t.Fatal(err)
	}
	templateFS := os.DirFS(filepath.Join(projectRoot, "templates"))
	if err := templates.Init(templateFS); err != nil {
		t.Fatal(err)
	}
}

// TestHandler is a helper function to test HTTP handlers
func TestHandler(t *testing.T, handler http.HandlerFunc, path string) {
	req := httptest.NewRequest("GET", path, nil)
	rr := httptest.NewRecorder()

	handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Error("expected non-empty response body, got empty string")
	}
}