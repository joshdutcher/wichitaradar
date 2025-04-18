package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"wichitaradar/internal/cache"
)

func TestHandleXML(t *testing.T) {
	// Create a temporary directory for test cache
	tempDir, err := os.MkdirTemp("", "xml_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create test XML file
	testXML := `<?xml version="1.0" encoding="UTF-8"?>
<graphicasts>
	<graphicast>
		<title>Test Story</title>
		<text>This is a test story</text>
	</graphicast>
</graphicasts>`

	xmlFile := filepath.Join(tempDir, "wxstory.xml")
	if err := os.WriteFile(xmlFile, []byte(testXML), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a new cache instance for testing, pointing to tempDir
	testCache := cache.New(tempDir, 15*time.Minute)

	// Create the handler instance using the factory and the test cache
	handler := NewXMLHandler(testCache)

	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   bool
	}{
		{
			name:       "missing path parameter",
			path:       "/xml",
			wantStatus: http.StatusBadRequest,
			wantBody:   false,
		},
		{
			name:       "valid path",
			path:       "/xml?path=ict/wxstory.xml",
			wantStatus: http.StatusOK,
			wantBody:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rr := httptest.NewRecorder()

			// Execute the handler instance
			handler(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

			if tt.wantBody && rr.Body.String() == "" {
				t.Error("expected non-empty response body, got empty string")
			}
		})
	}
}