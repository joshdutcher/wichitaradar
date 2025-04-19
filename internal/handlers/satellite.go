package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleSatellite handles the satellite page
func HandleSatellite(w http.ResponseWriter, r *http.Request) error {
	data := struct {
		Menu            *menu.Menu
		CurrentPath     string
		RefreshInterval int // Added for auto-refresh
	}{
		Menu:            menu.New(),
		CurrentPath:     r.URL.Path,
		RefreshInterval: 300, // 5 minutes
	}

	// Check if menu creation failed silently
	if data.Menu == nil {
		return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
	}

	// Get the specific template set for this page
	ts, err := templates.Get("satellite")
	if err != nil {
		return middleware.InternalError(fmt.Errorf("failed to get template set 'satellite': %w", err))
	}

	// Add a check to ensure ts is not nil, although Get should handle this
	if ts == nil {
		return middleware.InternalError(fmt.Errorf("got nil template set from templates.Get"))
	}

	// Execute the main template definition within this set
	if err := ts.ExecuteTemplate(w, "satellite", data); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to render template 'satellite': %w", err))
	}

	return nil
}
