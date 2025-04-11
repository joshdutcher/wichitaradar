package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

func HandleDisclaimer(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Menu        *menu.Menu
		CurrentPath string
	}{
		Menu:        menu.New(),
		CurrentPath: r.URL.Path,
	}

	if data.Menu == nil {
		http.Error(w, "menu.New() returned nil", http.StatusInternalServerError)
		return
	}

	ts, err := templates.Get("disclaimer")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get template set 'disclaimer': %v", err), http.StatusInternalServerError)
		return
	}

	if err := ts.ExecuteTemplate(w, "disclaimer", data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template 'disclaimer': %v", err), http.StatusInternalServerError)
		return
	}
}
