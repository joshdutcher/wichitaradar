package handlers

import (
	"net/http"
	"time"

	"wichitaradar/internal/cache"
)

var radarCache = cache.New("scraped/images", 5*time.Minute)

func HandleRadarImage(w http.ResponseWriter, r *http.Request) {
	// URL of the Weather Underground radar image
	url := "https://s.w-x.co/staticmaps/wu/wxtype/county_loc/sln/animate.png"

	// Get or download the cached image
	filepath, err := radarCache.GetFile(url, "radar.png", "https://www.wunderground.com/maps/radar/current/sln")
	if err != nil {
		http.Error(w, "Failed to get radar image", http.StatusInternalServerError)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filepath)
}
