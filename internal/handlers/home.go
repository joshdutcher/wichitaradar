package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleHome handles the home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "menu.New() returned nil", http.StatusInternalServerError)
		return
	}

	// Get the specific template set for this page
	ts, err := templates.Get("index") // Requesting "index" will get "index.page.html"
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get template set 'index': %v", err), http.StatusInternalServerError)
		return
	}

	// Execute the main template definition within this set (which should be "index")
	if err := ts.ExecuteTemplate(w, "index", data); err != nil {
		fmt.Fprintf(w, "Failed to render template 'index': %v", err) // Write error directly
		// Log the error server-side as well
		fmt.Printf("ERROR rendering template 'index': %v\n", err)
		return
	}
}
