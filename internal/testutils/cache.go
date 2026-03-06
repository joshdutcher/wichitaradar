package testutils

import (
	"fmt"
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

// MockErrorCacheProvider always returns an error from GetContent
type MockErrorCacheProvider struct{}

func (m *MockErrorCacheProvider) GetContent(url string, referer string, filename ...string) (io.ReadCloser, error) {
	return nil, fmt.Errorf("mock cache error")
}
