package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"wichitaradar/internal/handlers"
	"wichitaradar/internal/middleware"
	"wichitaradar/pkg/templates"
)

// mockCacheProvider implements cache.CacheProvider for testing
type mockCacheProvider struct {
	content string
}

func (m *mockCacheProvider) GetContent(url string, referer string, filename ...string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(m.content)), nil
}

func TestSetupServer(t *testing.T) {
	// Get the project root directory
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root directory: %v", err)
	}

	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "test_server")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create necessary directories
	for _, dir := range []string{
		filepath.Join(tempDir, "static"),
	} {
		if err := os.MkdirAll(dir, 0777); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Create a test static file
	staticFile := filepath.Join(tempDir, "static", "test.txt")
	if err := os.WriteFile(staticFile, []byte("test content"), 0666); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Initialize templates from the actual template files
	templateFS := os.DirFS(filepath.Join(projectRoot, "templates"))
	if err := templates.Init(templateFS); err != nil {
		t.Fatalf("Failed to initialize templates: %v", err)
	}

	// Create mock XML data
	mockXML := `<?xml version="1.0" encoding="UTF-8"?>
<xml>
	<graphicasts>
		<graphicast>
			<StartTime>0</StartTime>
			<EndTime>9999999999</EndTime>
			<radar>0</radar>
			<SmallImage>https://weather.gov/test.jpg</SmallImage>
			<description>Mock Story</description>
			<order>1</order>
		</graphicast>
	</graphicasts>
</xml>`

	// Create mock cache
	mockCache := &mockCacheProvider{content: mockXML}

	// Create a new mux for testing
	mux := http.NewServeMux()

	// Register routes for testing
	mux.Handle("/health", middleware.ErrorHandler(handlers.HandleHealth))
	mux.Handle("/outlook", middleware.ErrorHandler(handlers.HandleOutlook(mockCache)))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(tempDir, "static")))))

	// Add redirect route
	mux.Handle("/index.php", middleware.ErrorHandler(handlers.HandleRedirect("/")))

	// Add catch-all handler that handles both home page and 404s
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Not found")
			return
		}
		middleware.ErrorHandler(handlers.HandleHome).ServeHTTP(w, r)
	})

	// Test cases
	tests := []struct {
		name           string
		route          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "home route",
			route:          "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Wichita Radar",
		},
		{
			name:           "health check",
			route:          "/health",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "outlook route",
			route:          "/outlook",
			expectedStatus: http.StatusOK,
			expectedBody:   `src="https://weather.gov/test.jpg"`,
		},
		{
			name:           "static file",
			route:          "/static/test.txt",
			expectedStatus: http.StatusOK,
			expectedBody:   "test content",
		},
		{
			name:           "nonexistent route",
			route:          "/nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found",
		},
		{
			name:           "index.php redirect",
			route:          "/index.php",
			expectedStatus: http.StatusMovedPermanently,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.route, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("got status %d, want %d", w.Code, tt.expectedStatus)
			}

			if tt.expectedBody != "" && !bytes.Contains(w.Body.Bytes(), []byte(tt.expectedBody)) {
				t.Logf("Response body: %s", w.Body.String())
				t.Errorf("response body does not contain %q", tt.expectedBody)
			}
		})
	}
}

func TestMain(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "OK")
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// Test server is running
	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Server returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}
}
