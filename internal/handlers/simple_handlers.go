package handlers

import (
	"fmt"
	"net/http"

	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

// HandleSimplePage is a generic handler for simple template-based pages
func HandleSimplePage(templateName string) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		data := struct {
			Menu            *menu.Menu
			CurrentPath     string
			RefreshInterval int // Added for auto-refresh
		}{
			Menu:        menu.New(),
			CurrentPath: r.URL.Path,
			// Default to no refresh
			RefreshInterval: 0,
		}

		// Set refresh interval specifically for the watches page
		if templateName == "watches" {
			data.RefreshInterval = 600 // 10 minutes
		}

		if data.Menu == nil {
			return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
		}

		ts, err := templates.Get(templateName)
		if err != nil {
			return middleware.InternalError(fmt.Errorf("failed to get template set '%s': %w", templateName, err))
		}

		if err := ts.ExecuteTemplate(w, templateName, data); err != nil {
			return middleware.InternalError(fmt.Errorf("failed to render template '%s': %w", templateName, err))
		}

		return nil
	}
}

// HandleRedirect creates a handler that performs a 301 permanent redirect
func HandleRedirect(targetURL string) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		http.Redirect(w, r, targetURL, http.StatusMovedPermanently)
		return nil
	}
}
