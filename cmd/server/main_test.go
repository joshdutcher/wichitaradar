package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSetupServer(t *testing.T) {
	// Get the project root directory (two levels up from cmd/server)
	projectRoot := filepath.Join("..", "..")

	// Create a temporary directory for cache files
	tempDir, err := os.MkdirTemp("", "server-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test server setup with templates
	if err := setupServer(projectRoot, false); err != nil {
		t.Fatalf("setupServer failed: %v", err)
	}

	// Test that cache directories were created
	cacheDirs := []string{
		filepath.Join(projectRoot, "scraped/xml"),
		filepath.Join(projectRoot, "scraped/images"),
		filepath.Join(projectRoot, "scraped/temp"),
	}
	for _, dir := range cacheDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Cache directory %s was not created", dir)
		}
	}
	defer func() {
		// Clean up cache directories after test
		for _, dir := range cacheDirs {
			os.RemoveAll(dir)
		}
	}()

	// Test all registered routes
	routes := []struct {
		path       string
		wantStatus int
	}{
		{"/", http.StatusOK},
		{"/health", http.StatusOK},
		{"/outlook", http.StatusOK},
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

func TestMain(t *testing.T) {
	// Get the project root directory (two levels up from cmd/server)
	projectRoot := filepath.Join("..", "..")

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