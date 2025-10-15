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
