package cache

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Cache represents a local file cache with expiration
type Cache struct {
	cacheDir string
	cacheAge time.Duration
}

// New creates a new Cache instance
func New(cacheDir string, cacheAge time.Duration) *Cache {
	return &Cache{
		cacheDir: cacheDir,
		cacheAge: cacheAge,
	}
}

// GetCacheDir returns the cache directory path
func (c *Cache) GetCacheDir() string {
	return c.cacheDir
}

// Expired checks if a cached file has expired
func (c *Cache) Expired(filename string) bool {
	filepath := filepath.Join(c.cacheDir, filename)

	info, err := os.Stat(filepath)
	if err != nil {
		return true // File doesn't exist, consider it expired
	}

	return time.Since(info.ModTime()) > c.cacheAge
}

// DownloadFile downloads a file from a URL and saves it to the cache
func (c *Cache) DownloadFile(url, filename string, referer string) error {
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
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Log content length
	fmt.Printf("Content-Length header: %d\n", resp.ContentLength)

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
	defer file.Close()

	// Copy response body to file with progress tracking
	var bytesCopied int64
	bytesCopied, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	// Log actual bytes copied
	fmt.Printf("Bytes copied: %d\n", bytesCopied)

	// Get file info after writing
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	fmt.Printf("Cached file size: %d bytes\n", fileInfo.Size())

	// Verify we got all the content
	if resp.ContentLength > 0 && bytesCopied != resp.ContentLength {
		return fmt.Errorf("incomplete download: got %d bytes, expected %d", bytesCopied, resp.ContentLength)
	}

	return nil
}

// GetFile returns the path to a cached file, downloading it if necessary
func (c *Cache) GetFile(url, filename string, referer string) (string, error) {
	filepath := filepath.Join(c.cacheDir, filename)

	if c.Expired(filename) {
		if err := c.DownloadFile(url, filename, referer); err != nil {
			return "", err
		}
	}

	return filepath, nil
}
