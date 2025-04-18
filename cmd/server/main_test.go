package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
	// No longer need direct cache/handlers import here for mocking
)

func TestSetupServer(t *testing.T) {
	// Get the project root directory (two levels up from cmd/server)
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatalf("Failed to get project root directory: %v", err)
	}

	// --- Mocking Prerequisite Start ---
	// Define paths for cache directories that setupServer will create
	xmlCacheDir := filepath.Join(projectRoot, "cache/xml")
	imagesCacheDir := filepath.Join(projectRoot, "cache/images")
	animatedCacheDir := filepath.Join(projectRoot, "cache/animated")
	cacheDirsToClean := []string{xmlCacheDir, imagesCacheDir, animatedCacheDir}

	// Ensure the directories exist *before* setupServer is called
	// This is needed because setupServer expects to create them,
	// and we need the xml dir to exist to place the mock file.
	for _, dir := range cacheDirsToClean {
		if err := os.MkdirAll(dir, 0777); err != nil {
			t.Fatalf("Failed to pre-create test cache directory %s: %v", dir, err)
		}
	}
	// Cleanup these directories after the test
	defer func() {
		for _, dir := range cacheDirsToClean {
			os.RemoveAll(dir)
		}
	}()

	// Create a mock wxstory.xml file in the *real* cache path *before* setupServer runs
	testXMLContent := `<?xml version="1.0" encoding="UTF-8"?><graphicasts><graphicast><title>Mock Story From Main Test</title><text>...</text></graphicast></graphicasts>`
	mockXMLPath := filepath.Join(xmlCacheDir, "wxstory.xml")
	if err := os.WriteFile(mockXMLPath, []byte(testXMLContent), 0644); err != nil {
		t.Fatalf("Failed to create mock XML file: %v", err)
	}
	// --- Mocking Prerequisite End ---

	// Test server setup with templates.
	// setupServer will now create its own cache instance pointing to the directory
	// containing our mock file.
	if err := setupServer(projectRoot, false); err != nil {
		t.Fatalf("setupServer failed: %v", err)
	}

	// Test all registered routes
	routes := []struct {
		path       string
		wantStatus int
	}{
		{"/", http.StatusOK},
		{"/health", http.StatusOK},
		{"/outlook", http.StatusOK}, // Should hit the handler using the mock file
		{"/satellite", http.StatusOK},
		{"/watches", http.StatusOK},
		{"/temperatures", http.StatusOK},
		{"/rainfall", http.StatusOK},
		{"/resources", http.StatusOK},
		{"/about", http.StatusOK},
		{"/disclaimer", http.StatusOK},
		{"/donate", http.StatusOK},
		{"/static/nonexistent.css", http.StatusNotFound},
	}

	for _, route := range routes {
		t.Run(route.path, func(t *testing.T) {
			req := httptest.NewRequest("GET", route.path, nil)
			rr := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rr, req)

			if status := rr.Code; status != route.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, route.wantStatus)
			}
		})
	}

	// Test static file handling
	t.Run("static file handling", func(t *testing.T) {
		// Create a test static file
		testFile := filepath.Join(projectRoot, "static", "test.txt")
		if err := os.MkdirAll(filepath.Dir(testFile), 0755); err != nil {
			t.Fatalf("Failed to create static directory: %v", err)
		}
		if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(testFile)

		req := httptest.NewRequest("GET", "/static/test.txt", nil)
		rr := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("static file returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if body := rr.Body.String(); body != "test content" {
			t.Errorf("static file returned wrong content: got %q want %q",
				body, "test content")
		}
	})
}

// TestMain remains the same as it doesn't directly test the XML handler logic
func TestMain(t *testing.T) {
	// Get the project root directory (two levels up from cmd/server)
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}

	// Save original environment variables
	originalPort := os.Getenv("PORT")

	// Set up test environment
	os.Setenv("PORT", "8080")

	// Save current working directory
	originalWorkDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// Change to project root directory
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.Chdir(originalWorkDir)
	}()

	// Test main function
	go func() {
		// Reset the default ServeMux to avoid conflicts
		http.DefaultServeMux = http.NewServeMux()
		main()
	}()

	// Wait for server to start
	time.Sleep(100 * time.Millisecond)

	// Test server is running
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Server returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}

	// Restore environment
	os.Setenv("PORT", originalPort)
}