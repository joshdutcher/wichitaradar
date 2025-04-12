package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"wichitaradar/internal/handlers"
	"wichitaradar/pkg/templates"
)

func main() {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	// Create cache directories with proper permissions
	cacheDirs := []string{
		filepath.Join(workDir, "scraped/xml"),
		filepath.Join(workDir, "scraped/images"),
		filepath.Join(workDir, "scraped/temp"),
	}
	for _, dir := range cacheDirs {
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Printf("Failed to create cache directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}

	// Initialize templates from the "templates" directory
	templateFS := os.DirFS(filepath.Join(workDir, "templates"))
	if err := templates.Init(templateFS); err != nil {
		fmt.Printf("Failed to initialize templates: %v\n", err)
		os.Exit(1)
	}

	// Serve static files
	staticDir := filepath.Join(workDir, "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Register handlers
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/outlook", handlers.HandleOutlook)
	http.HandleFunc("/satellite", handlers.HandleSatellite)
	http.HandleFunc("/watches", handlers.HandleWatches)
	http.HandleFunc("/temperatures", handlers.HandleTemperatures)
	http.HandleFunc("/rainfall", handlers.HandleRainfall)
	http.HandleFunc("/resources", handlers.HandleResources)
	http.HandleFunc("/about", handlers.HandleAbout)
	http.HandleFunc("/disclaimer", handlers.HandleDisclaimer)
	http.HandleFunc("/donate", handlers.HandleDonate)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Start server
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
