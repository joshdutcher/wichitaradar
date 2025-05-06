package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"wichitaradar/internal/cache"
	"wichitaradar/internal/handlers"
	"wichitaradar/internal/middleware"
	"wichitaradar/pkg/templates"

	"github.com/getsentry/sentry-go"
)

func init() {
	// Initialize Sentry
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		log.Printf("Warning: SENTRY_DSN not set, Sentry will not be initialized")
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		Environment: func() string {
			if os.Getenv("ENV") == "production" {
				return "production"
			}
			return "development"
		}(),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

// setupServer configures the HTTP server with all routes and middleware
func setupServer(workDir string, skipTemplates bool) error {
	// Initialize caches
	projectRoot, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Failed to get project root directory: %v", err)
	}

	// Create cache directories
	cacheDirs := []string{
		cache.GetXMLCacheDir(projectRoot),
		cache.GetAnimatedCacheDir(projectRoot),
	}

	for _, dir := range cacheDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create cache directory %s: %v", dir, err)
		}
	}

	// Create Cache instances
	xmlCache := cache.NewFileCache(cache.GetXMLCacheDir(projectRoot), 5*time.Minute)

	// Initialize templates from the "templates" directory
	if !skipTemplates {
		templateFS := os.DirFS(filepath.Join(workDir, "templates"))
		if err := templates.Init(templateFS); err != nil {
			return fmt.Errorf("failed to initialize templates: %v", err)
		}
	}

	// Create a new mux to wrap with middleware
	mux := http.NewServeMux()

	// Serve static files
	staticDir := filepath.Join(workDir, "static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Register routes
	mux.Handle("/", middleware.ErrorHandler(handlers.HandleHome))
	mux.Handle("/health", middleware.ErrorHandler(handlers.HandleHealth))
	mux.Handle("/temperatures", middleware.ErrorHandler(handlers.HandleTemperatures))
	mux.Handle("/rainfall", middleware.ErrorHandler(handlers.HandleRainfall))
	mux.Handle("/satellite", middleware.ErrorHandler(handlers.HandleSatellite))
	mux.Handle("/outlook", middleware.ErrorHandler(handlers.HandleOutlook(xmlCache)))
	mux.Handle("/watches", middleware.ErrorHandler(handlers.HandleSimplePage("watches")))
	mux.Handle("/resources", middleware.ErrorHandler(handlers.HandleSimplePage("resources")))
	mux.Handle("/about", middleware.ErrorHandler(handlers.HandleSimplePage("about")))
	mux.Handle("/disclaimer", middleware.ErrorHandler(handlers.HandleSimplePage("disclaimer")))
	mux.Handle("/donate", middleware.ErrorHandler(handlers.HandleSimplePage("donate")))
	mux.Handle("/api/image-error", middleware.ErrorHandler(handlers.HandleImageError))

	// Set the mux as the default handler for all routes
	http.Handle("/", mux)

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
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		os.Exit(1)
	}
}
