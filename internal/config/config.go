package config

import "os"

// GetWeatherServiceURL returns the URL for the weather service
func GetWeatherServiceURL() string {
	// Default to localhost for development
	url := os.Getenv("WEATHER_SERVICE_URL")
	if url == "" {
		url = "http://localhost:316"
	}
	return url
}
