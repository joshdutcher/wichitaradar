package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleHome handles the home page
func HandleHome(w http.ResponseWriter, r *http.Request) error {
	// Create template data
	data := struct {
		Menu            *menu.Menu
		CurrentPath     string
		RefreshInterval int // Added for auto-refresh
	}{
		Menu:            menu.New(),
		CurrentPath:     r.URL.Path,
		RefreshInterval: 300,
	}

	// Check if menu creation failed silently
	if data.Menu == nil {
		return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
	}

	// Get the specific template set for this page
	ts, err := templates.Get("index") // Requesting "index" will get "index.page.html"
	if err != nil {
		return middleware.InternalError(fmt.Errorf("failed to get template set 'index': %w", err))
	}

	// Execute the main template definition within this set (which should be "index")
	if err := ts.ExecuteTemplate(w, "index", data); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to render template 'index': %w", err))
	}

	return nil
}
