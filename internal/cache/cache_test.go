package cache

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cacheDir := "test_cache"
	cacheAge := 5 * time.Minute

	c := New(cacheDir, cacheAge)

	if c.cacheDir != cacheDir {
		t.Errorf("expected cacheDir %q, got %q", cacheDir, c.cacheDir)
	}

	if c.cacheAge != cacheAge {
		t.Errorf("expected cacheAge %v, got %v", cacheAge, c.cacheAge)
	}
}

func TestGetCacheDir(t *testing.T) {
	cacheDir := "test_cache"
	c := New(cacheDir, 5*time.Minute)

	if got := c.GetCacheDir(); got != cacheDir {
		t.Errorf("expected %q, got %q", cacheDir, got)
	}
}

func TestExpired(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "cache_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test file
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	c := New(tempDir, 1*time.Second)

	// Test with non-existent file
	if !c.Expired("nonexistent.txt") {
		t.Error("expected expired for non-existent file")
	}

	// Test with existing file that hasn't expired
	if c.Expired("test.txt") {
		t.Error("expected not expired for new file")
	}

	// Wait for file to expire
	time.Sleep(2 * time.Second)
	if !c.Expired("test.txt") {
		t.Error("expected expired after waiting")
	}
}

func TestGetCacheDirs(t *testing.T) {
	workDir := "/test/work/dir"
	expected := []string{
		"/test/work/dir/scraped/xml",
		"/test/work/dir/scraped/images",
		"/test/work/dir/scraped/temp",
	}

	got := GetCacheDirs(workDir)
	if len(got) != len(expected) {
		t.Errorf("expected %d directories, got %d", len(expected), len(got))
	}

	for i, dir := range got {
		if dir != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], dir)
		}
	}
}

func TestGetXMLCacheDir(t *testing.T) {
	workDir := "/test/work/dir"
	expected := "/test/work/dir/scraped/xml"

	if got := GetXMLCacheDir(workDir); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestGetImagesCacheDir(t *testing.T) {
	workDir := "/test/work/dir"
	expected := "/test/work/dir/scraped/images"

	if got := GetImagesCacheDir(workDir); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestGetTempCacheDir(t *testing.T) {
	workDir := "/test/work/dir"
	expected := "/test/work/dir/scraped/temp"

	if got := GetTempCacheDir(workDir); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestDownloadFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "cache_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("User-Agent") == "" {
			t.Error("expected User-Agent header")
		}
		if referer := r.Header.Get("Referer"); referer != "http://test-referer" {
			t.Errorf("expected Referer header %q, got %q", "http://test-referer", referer)
		}

		// Return test content
		w.Write([]byte("test content"))
	}))
	defer ts.Close()

	c := New(tempDir, 5*time.Minute)

	// Test successful download
	err = c.DownloadFile(ts.URL, "test.txt", "http://test-referer")
	if err != nil {
		t.Errorf("DownloadFile() error = %v", err)
	}

	// Verify file was created
	content, err := os.ReadFile(filepath.Join(tempDir, "test.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "test content" {
		t.Errorf("expected content %q, got %q", "test content", string(content))
	}

	// Test invalid URL
	err = c.DownloadFile("invalid-url", "test.txt", "")
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}

func TestGetFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "cache_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test content"))
	}))
	defer ts.Close()

	c := New(tempDir, 5*time.Minute)

	// Test getting a file that doesn't exist (should trigger download)
	filepath, err := c.GetFile(ts.URL, "test.txt", "")
	if err != nil {
		t.Errorf("GetFile() error = %v", err)
	}

	// Verify file was downloaded
	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "test content" {
		t.Errorf("expected content %q, got %q", "test content", string(content))
	}

	// Test getting the same file again (should use cache)
	filepath2, err := c.GetFile(ts.URL, "test.txt", "")
	if err != nil {
		t.Errorf("GetFile() error = %v", err)
	}
	if filepath2 != filepath {
		t.Errorf("expected same filepath, got different ones: %q != %q", filepath, filepath2)
	}
}