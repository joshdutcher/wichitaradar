package handlers

import (
	"fmt"
	"net/http"
	"time"

	"wichitaradar/menu"
	"wichitaradar/pkg/templates"
)

func HandleRainfall(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Menu        *menu.Menu
		CurrentPath string
		CurrentDate string
	}{
		Menu:        menu.New(),
		CurrentPath: r.URL.Path,
		CurrentDate: time.Now().Format("20060102"),
	}

	if data.Menu == nil {
		http.Error(w, "menu.New() returned nil", http.StatusInternalServerError)
		return
	}

	ts, err := templates.Get("rainfall")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get template set 'rainfall': %v", err), http.StatusInternalServerError)
		return
	}

	if ts == nil {
		http.Error(w, "Got nil template set from templates.Get", http.StatusInternalServerError)
		return
	}

	if err := ts.ExecuteTemplate(w, "rainfall", data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template 'rainfall': %v", err), http.StatusInternalServerError)
		return
	}
}
