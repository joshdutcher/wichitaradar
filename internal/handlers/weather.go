package handlers

import (
	"net/http"
	"path/filepath"
	"time"

	"wichitaradar/internal/cache"
)

var (
	// Cache for weather.gov images (15 minutes)
	weatherCache = cache.New("scraped/images", 15*time.Minute)

	// Cache for SPC outlook images (1 hour)
	spcCache = cache.New("scraped/images", time.Hour)
)

// HandleWeatherImage handles images from weather.gov
func HandleWeatherImage(w http.ResponseWriter, r *http.Request) {
	// Extract the image path from the URL
	imagePath := r.URL.Query().Get("path")
	if imagePath == "" {
		http.Error(w, "Missing image path", http.StatusBadRequest)
		return
	}

	// Construct the full URL
	url := "http://weather.gov/" + imagePath

	// Get or download the cached image
	filename := filepath.Base(imagePath)
	filepath, err := weatherCache.GetFile(url, filename, "http://www.weather.gov/ict/")
	if err != nil {
		http.Error(w, "Failed to get weather image", http.StatusInternalServerError)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filepath)
}

// HandleSPCImage handles images from the Storm Prediction Center
func HandleSPCImage(w http.ResponseWriter, r *http.Request) {
	// Extract the image type from the URL
	imageType := r.URL.Query().Get("type")
	if imageType == "" {
		http.Error(w, "Missing image type", http.StatusBadRequest)
		return
	}

	// Construct the full URL
	url := "http://www.spc.noaa.gov/products/outlook/" + imageType + "otlk.gif"

	// Get or download the cached image
	filepath, err := spcCache.GetFile(url, imageType+".gif", "http://www.spc.noaa.gov/")
	if err != nil {
		http.Error(w, "Failed to get SPC image", http.StatusInternalServerError)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filepath)
}
