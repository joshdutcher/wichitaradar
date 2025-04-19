package handlers

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"log"
	"wichitaradar/internal/cache"
	"wichitaradar/internal/middleware"
)

// Cache for animated images (5 minutes)
var animatedCache *cache.Cache

func init() {
	// Get the project root directory
	projectRoot, err := filepath.Abs("../..")
	if err != nil {
		log.Fatalf("Failed to get project root directory: %v", err)
	}
	animatedCache = cache.NewFileCache(cache.GetAnimatedCacheDir(projectRoot), 5*time.Minute)
}

// HandleWundergroundAnimatedRadar handles requests for the Wunderground animated radar image
func HandleWundergroundAnimatedRadar(cacheProvider cache.CacheProvider) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Get the image content from cache
		url := "https://s.w-x.co/staticmaps/wu/wxtype/county_loc/sln/animate.png"
		content, err := cacheProvider.GetContent(url, "https://www.wunderground.com/", "wunderground-animated-radar.png")
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to fetch animated image: %w", err))
		}
		defer content.Close()

		// Set content type and serve the image
		w.Header().Set("Content-Type", "image/png")
		_, err = io.Copy(w, content)
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to serve animated image: %w", err))
		}
		return nil
	}
}