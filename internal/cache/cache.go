package cache

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// CacheProvider defines the interface for cache operations
type CacheProvider interface {
	// GetContent returns the content for a given URL, using the cache if available
	GetContent(url string, referer string, filename ...string) (io.ReadCloser, error)
}

// Cache represents a local file cache with expiration
type Cache struct {
	cacheDir string
	cacheAge time.Duration
}

// NewFileCache creates a new Cache instance
func NewFileCache(cacheDir string, cacheAge time.Duration) *Cache {
	return &Cache{
		cacheDir: cacheDir,
		cacheAge: cacheAge,
	}
}

func (c *Cache) isExpired(filename string) bool {
	filepath := filepath.Join(c.cacheDir, filename)

	info, err := os.Stat(filepath)
	if err != nil {
		return true // File doesn't exist, consider it expired
	}

	return time.Since(info.ModTime()) > c.cacheAge
}

func (c *Cache) downloadToCache(url, filename string, referer string) error {
	// Create HTTP client with custom headers
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.13) Gecko/20080311 Firefox/2.0.0.13")
	if referer != "" {
		req.Header.Set("Referer", referer)
	}

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(c.cacheDir, 0755); err != nil {
		return err
	}

	// Create file
	filepath := filepath.Join(c.cacheDir, filename)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	// Copy response body to file
	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}

// GetContent implements CacheProvider
func (c *Cache) GetContent(url string, referer string, filename ...string) (io.ReadCloser, error) {
	// Use provided filename or generate from URL
	var cacheFilename string
	if len(filename) > 0 {
		cacheFilename = filename[0]
	} else {
		cacheFilename = filepath.Base(url)
	}

	// Check if we need to download
	if c.isExpired(cacheFilename) {
		if err := c.downloadToCache(url, cacheFilename, referer); err != nil {
			return nil, fmt.Errorf("failed to download: %w", err)
		}
	}

	// Open and return the file
	filepath := filepath.Join(c.cacheDir, cacheFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache file: %w", err)
	}

	return file, nil
}
