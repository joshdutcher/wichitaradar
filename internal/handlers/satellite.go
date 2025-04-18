package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleSatellite handles the satellite page
func HandleSatellite(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "menu.New() returned nil", http.StatusInternalServerError)
		return
	}

	// Get the specific template set for this page
	ts, err := templates.Get("satellite")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get template set 'satellite': %v", err), http.StatusInternalServerError)
		return
	}

	// Add a check to ensure ts is not nil, although Get should handle this
	if ts == nil {
		http.Error(w, "Got nil template set from templates.Get", http.StatusInternalServerError)
		return
	}
	// Execute the main template definition within this set
	if err := ts.ExecuteTemplate(w, "satellite", data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template 'satellite': %v", err), http.StatusInternalServerError)
		return
	}
}
