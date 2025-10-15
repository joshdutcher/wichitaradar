package testutils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
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

	body := rr.Body.String()
	if body == "" {
		t.Error("expected non-empty response body, got empty string")
	}

	// Check for RefreshInterval in the response body based on the page
	switch path {
	case "/":
		if !strings.Contains(body, `<meta http-equiv="refresh" content="300" />`) {
			t.Error("home page should have refreshInterval set to 300")
		}
	case "/satellite":
		if !strings.Contains(body, `<meta http-equiv="refresh" content="300" />`) {
			t.Error("satellite page should have refreshInterval set to 300")
		}
	case "/watches":
		if !strings.Contains(body, `<meta http-equiv="refresh" content="600" />`) {
			t.Error("watches page should have refreshInterval set to 600")
		}
	case "/outlook":
		if !strings.Contains(body, `<meta http-equiv="refresh" content="1800" />`) {
			t.Error("outlook page should have refreshInterval set to 1800")
		}
	case "/about", "/disclaimer", "/donate", "/resources":
		if strings.Contains(body, `<meta http-equiv="refresh"`) {
			t.Errorf("%s page should not have a refresh meta tag", path)
		}
	}
}
