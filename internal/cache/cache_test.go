package cache

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		cacheDir    string
		cacheAge    time.Duration
		description string
	}{
		{
			name:        "valid cache configuration",
			cacheDir:    "test_cache",
			cacheAge:    5 * time.Minute,
			description: "should create cache with specified directory and age",
		},
		{
			name:        "zero cache age",
			cacheDir:    "test_cache",
			cacheAge:    0,
			description: "should create cache with zero age",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewFileCache(tt.cacheDir, tt.cacheAge)

			if c.cacheDir != tt.cacheDir {
				t.Errorf("%s: expected cacheDir %q, got %q", tt.description, tt.cacheDir, c.cacheDir)
			}

			if c.cacheAge != tt.cacheAge {
				t.Errorf("%s: expected cacheAge %v, got %v", tt.description, tt.cacheAge, c.cacheAge)
			}
		})
	}
}

func TestGetContent(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(string) string
		cacheAge    time.Duration
		waitTime    time.Duration
		url         string
		referer     string
		filename    string
		want        string
		description string
	}{
		{
			name: "non-existent file",
			setup: func(dir string) string {
				return "nonexistent.txt"
			},
			cacheAge:    time.Minute,
			waitTime:    0,
			url:         "http://example.com/test.txt",
			referer:     "http://example.com/",
			filename:    "test.txt",
			want:        "test content",
			description: "should download and cache new file",
		},
		{
			name: "fresh file",
			setup: func(dir string) string {
				filename := filepath.Join(dir, "fresh.txt")
				os.WriteFile(filename, []byte("test content"), 0644)
				return "fresh.txt"
			},
			cacheAge:    time.Minute,
			waitTime:    0,
			url:         "http://example.com/fresh.txt",
			referer:     "http://example.com/",
			filename:    "fresh.txt",
			want:        "test content",
			description: "should return cached file",
		},
		{
			name: "expired file",
			setup: func(dir string) string {
				filename := filepath.Join(dir, "expired.txt")
				os.WriteFile(filename, []byte("old content"), 0644)
				return "expired.txt"
			},
			cacheAge:    time.Second,
			waitTime:    2 * time.Second,
			url:         "http://example.com/expired.txt",
			referer:     "http://example.com/",
			filename:    "expired.txt",
			want:        "new content",
			description: "should download new content for expired file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary directory for testing
			tempDir, err := os.MkdirTemp("", "cache_test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tempDir)

			// Setup test file
			filename := tt.setup(tempDir)

			// Create cache with test configuration
			c := NewFileCache(tempDir, tt.cacheAge)

			// Create a test server to serve content
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tt.want))
			}))
			defer ts.Close()

			// Wait if specified
			if tt.waitTime > 0 {
				time.Sleep(tt.waitTime)
			}

			// Get content
			content, err := c.GetContent(ts.URL+"/"+filename, tt.referer, tt.filename)
			if err != nil {
				t.Fatalf("%s: GetContent() error = %v", tt.description, err)
			}
			defer content.Close()

			// Read content
			data, err := io.ReadAll(content)
			if err != nil {
				t.Fatalf("%s: ReadAll() error = %v", tt.description, err)
			}

			// Check content
			if string(data) != tt.want {
				t.Errorf("%s: expected %q, got %q", tt.description, tt.want, string(data))
			}
		})
	}
}

func TestGetCacheDirs(t *testing.T) {
	tests := []struct {
		name        string
		workDir     string
		want        []string
		description string
	}{
		{
			name:    "absolute path",
			workDir: "/test/work/dir",
			want: []string{
				"/test/work/dir/scraped/xml",
				"/test/work/dir/scraped/images",
				"/test/work/dir/scraped/temp",
			},
			description: "should return absolute paths for cache directories",
		},
		{
			name:    "relative path",
			workDir: "work/dir",
			want: []string{
				"work/dir/scraped/xml",
				"work/dir/scraped/images",
				"work/dir/scraped/temp",
			},
			description: "should return relative paths for cache directories",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCacheDirs(tt.workDir)
			if len(got) != len(tt.want) {
				t.Errorf("%s: expected %d directories, got %d", tt.description, len(tt.want), len(got))
			}

			for i, dir := range got {
				if dir != tt.want[i] {
					t.Errorf("%s: expected %q, got %q", tt.description, tt.want[i], dir)
				}
			}
		})
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