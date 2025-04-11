package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleHome handles the home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Home handler reached!") // Simple response for testing
	// Temporarily commented out template logic
	// Create template data
	data := struct {
		Menu        *menu.Menu
		CurrentPath string
	}{
		Menu:        menu.New(), // This might cause issues if menu package isn't imported
		CurrentPath: r.URL.Path,
	}

	// Check if menu creation failed silently (though unlikely)
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

	// Add a check to ensure ts is not nil, although Get should handle this
	if ts == nil {
		http.Error(w, "Got nil template set from templates.Get", http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "Successfully got template set for page: %s!", data.Title) // REMOVED test print

	// Execute the main template definition within this set (which should be "index")
	if err := ts.ExecuteTemplate(w, "index", data); err != nil {
		fmt.Fprintf(w, "Failed to render template 'index': %v", err) // Write error directly
		// Log the error server-side as well
		fmt.Printf("ERROR rendering template 'index': %v\n", err)
		return
	}
}
