package config

import (
	"os"
	"testing"
)

func TestGetWeatherServiceURL(t *testing.T) {
	// Save original environment variable
	originalURL := os.Getenv("WEATHER_SERVICE_URL")
	defer os.Setenv("WEATHER_SERVICE_URL", originalURL)

	// Test case 1: Environment variable set
	expectedURL := "http://test-service:8080"
	os.Setenv("WEATHER_SERVICE_URL", expectedURL)
	if got := GetWeatherServiceURL(); got != expectedURL {
		t.Errorf("expected %q, got %q", expectedURL, got)
	}

	// Test case 2: Environment variable not set
	os.Unsetenv("WEATHER_SERVICE_URL")
	expectedDefault := "http://localhost:316"
	if got := GetWeatherServiceURL(); got != expectedDefault {
		t.Errorf("expected default %q, got %q", expectedDefault, got)
	}
}