package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleSimplePage is a generic handler for simple template-based pages
func HandleSimplePage(templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		ts, err := templates.Get(templateName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get template set '%s': %v", templateName, err), http.StatusInternalServerError)
			return
		}

		if err := ts.ExecuteTemplate(w, templateName, data); err != nil {
			http.Error(w, fmt.Sprintf("Failed to render template '%s': %v", templateName, err), http.StatusInternalServerError)
			return
		}
	}
}