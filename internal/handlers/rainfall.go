package handlers

import (
	"fmt"
	"net/http"
	"time"

	"wichitaradar/internal/middleware"
	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

func HandleRainfall(w http.ResponseWriter, r *http.Request) error {
	data := struct {
		Menu            *menu.Menu
		CurrentPath     string
		CurrentDate     string
		RefreshInterval int
	}{
		Menu:            menu.New(),
		CurrentPath:     r.URL.Path,
		CurrentDate:     time.Now().Format("20060102"),
		RefreshInterval: 0, // No auto-refresh for rainfall page
	}

	if data.Menu == nil {
		return middleware.InternalError(fmt.Errorf("menu.New() returned nil"))
	}

	ts, err := templates.Get("rainfall")
	if err != nil {
		return middleware.InternalError(fmt.Errorf("failed to get template set 'rainfall': %w", err))
	}

	if ts == nil {
		return middleware.InternalError(fmt.Errorf("got nil template set from templates.Get"))
	}

	if err := ts.ExecuteTemplate(w, "rainfall", data); err != nil {
		return middleware.InternalError(fmt.Errorf("failed to render template 'rainfall': %w", err))
	}

	return nil
}
