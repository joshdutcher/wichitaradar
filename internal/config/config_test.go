package config

import (
	"os"
	"testing"
)

func TestGetWeatherServiceURL(t *testing.T) {
	tests := []struct {
		name        string
		envValue    *string // nil means unset
		expected    string
		description string
	}{
		{
			name:        "custom URL from environment",
			envValue:    stringPtr("http://test-service:8080"),
			expected:    "http://test-service:8080",
			description: "should return custom URL when environment variable is set",
		},
		{
			name:        "default URL when env not set",
			envValue:    nil,
			expected:    "http://localhost:316",
			description: "should return default URL when environment variable is not set",
		},
		{
			name:        "empty URL in environment",
			envValue:    stringPtr(""),
			expected:    "http://localhost:316",
			description: "should return default URL when environment variable is empty",
		},
	}

	// Save original environment variable
	originalURL := os.Getenv("WEATHER_SERVICE_URL")
	defer os.Setenv("WEATHER_SERVICE_URL", originalURL)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment
			if tt.envValue == nil {
				os.Unsetenv("WEATHER_SERVICE_URL")
			} else {
				os.Setenv("WEATHER_SERVICE_URL", *tt.envValue)
			}

			// Test
			got := GetWeatherServiceURL()
			if got != tt.expected {
				t.Errorf("%s: got %q, want %q", tt.description, got, tt.expected)
			}
		})
	}
}

// Helper function to get pointer to string
func stringPtr(s string) *string {
	return &s
}
