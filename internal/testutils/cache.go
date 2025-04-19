package testutils

import (
	"io"
	"strings"
)

// MockCacheProvider implements CacheProvider for testing
type MockCacheProvider struct {
	Content string
}

func (m *MockCacheProvider) GetContent(url string, referer string, filename ...string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(m.Content)), nil
}

// mockCacheProvider implements cache.CacheProvider for testing
type mockCacheProvider struct {
	content string
}

func (m *mockCacheProvider) GetContent(url string, referer string, filename ...string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(m.content)), nil
}