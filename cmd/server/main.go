package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"wichitaradar/internal/cache"
	"wichitaradar/internal/handlers"
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

	// Register routes
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/outlook", handlers.HandleOutlook)
	http.HandleFunc("/satellite", handlers.HandleSatellite)
	http.HandleFunc("/watches", handlers.HandleSimplePage("watches"))
	http.HandleFunc("/temperatures", handlers.HandleTemperatures)
	http.HandleFunc("/rainfall", handlers.HandleRainfall)
	http.HandleFunc("/resources", handlers.HandleSimplePage("resources"))
	http.HandleFunc("/about", handlers.HandleSimplePage("about"))
	http.HandleFunc("/disclaimer", handlers.HandleSimplePage("disclaimer"))
	http.HandleFunc("/donate", handlers.HandleSimplePage("donate"))
	http.HandleFunc("/api/image-error", handlers.HandleImageError)

	return nil
}

func main() {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	// Setup server
	if err := setupServer(workDir, false); err != nil {
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
