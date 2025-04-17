package handlers

import (
	"net/http"
	"path/filepath"
	"time"

	"wichitaradar/internal/cache"
)

// Cache for animated images (5 minutes)
var animatedCache *cache.Cache

func init() {
	// Get the project root directory
	projectRoot, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	animatedCache = cache.New(cache.GetImagesCacheDir(projectRoot), 5*time.Minute)
}

// HandleWundergroundAnimatedRadar handles requests for the Wunderground animated radar image
func HandleWundergroundAnimatedRadar(w http.ResponseWriter, r *http.Request) {
	// Get the cached image file
	cacheFile := filepath.Join(animatedCache.GetCacheDir(), "wunderground-animated-radar.png")
	if animatedCache.Expired("wunderground-animated-radar.png") {
		// Download fresh image
		url := "https://s.w-x.co/staticmaps/wu/wxtype/county_loc/sln/animate.png"
		if err := animatedCache.DownloadFile(url, "wunderground-animated-radar.png", "https://www.wunderground.com/"); err != nil {
			http.Error(w, "Failed to fetch animated image", http.StatusInternalServerError)
			return
		}
	}

	// Read and serve the cached file
	http.ServeFile(w, r, cacheFile)
}