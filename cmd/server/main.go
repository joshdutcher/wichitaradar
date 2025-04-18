package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"wichitaradar/internal/cache"
	"wichitaradar/internal/handlers"
	"wichitaradar/internal/middleware"
	"wichitaradar/pkg/templates"
)

// setupServer configures the HTTP server with all routes and middleware
func setupServer(workDir string, skipTemplates bool) error {
	// Create cache directories with proper permissions
	cacheDirs := cache.GetCacheDirs(workDir)
	for _, dir := range cacheDirs {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return fmt.Errorf("failed to create cache directory %s: %v", dir, err)
		}
	}

	// Create Cache instances
	// Adjust cache durations as needed
	xmlCache := cache.New(cache.GetXMLCacheDir(workDir), 15*time.Minute)
	// animatedCache := cache.New(cache.GetAnimatedCacheDir(workDir), 5*time.Minute) // Add if needed by other handlers
	// imagesCache := cache.New(cache.GetImagesCacheDir(workDir), 60*time.Minute) // Add if needed by other handlers

	// Initialize templates from the "templates" directory
	if !skipTemplates {
		templateFS := os.DirFS(filepath.Join(workDir, "templates"))
		if err := templates.Init(templateFS); err != nil {
			return fmt.Errorf("failed to initialize templates: %v", err)
		}
	}

	// Serve static files
	staticDir := filepath.Join(workDir, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Create a new mux to wrap with middleware
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", handlers.HandleHome)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/outlook", handlers.HandleOutlook(xmlCache))
	mux.HandleFunc("/satellite", handlers.HandleSatellite)
	mux.HandleFunc("/watches", handlers.HandleSimplePage("watches"))
	mux.HandleFunc("/temperatures", handlers.HandleTemperatures)
	mux.HandleFunc("/rainfall", handlers.HandleRainfall)
	mux.HandleFunc("/resources", handlers.HandleSimplePage("resources"))
	mux.HandleFunc("/about", handlers.HandleSimplePage("about"))
	mux.HandleFunc("/disclaimer", handlers.HandleSimplePage("disclaimer"))
	mux.HandleFunc("/donate", handlers.HandleSimplePage("donate"))
	mux.HandleFunc("/api/image-error", handlers.HandleImageError)
	mux.HandleFunc("/api/wunderground/animated-radar", handlers.HandleWundergroundAnimatedRadar)

	// Register XML handler using the new factory function
	mux.HandleFunc("/xml", handlers.HandleXML(xmlCache))

	// Wrap the mux with error handling middleware
	http.Handle("/", middleware.ErrorHandler(mux))

	return nil
}

func main() {
	// Get the project root directory
	projectRoot, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("Failed to get project root directory: %v\n", err)
		os.Exit(1)
	}

	// Setup server
	if err := setupServer(projectRoot, false); err != nil {
		fmt.Printf("Failed to setup server: %v\n", err)
		os.Exit(1)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Start server
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
