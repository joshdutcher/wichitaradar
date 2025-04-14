package menu

import (
	"path/filepath"
	"time"
)

// MenuItem represents a single item in the navigation menu
type MenuItem struct {
	Label   string
	URL     string
	Tooltip string
	IsNew   bool
}

// Menu represents the entire navigation menu
type Menu struct {
	Items    []MenuItem
	LoadTime time.Time
}

// MenuProvider defines the interface for creating menus
type MenuProvider interface {
	New() *Menu
}

// DefaultMenuProvider is the default implementation of MenuProvider
type DefaultMenuProvider struct{}

// New creates a new menu using the default provider
func New() *Menu {
	return DefaultMenuProvider{}.New()
}

// New creates a new menu
func (p DefaultMenuProvider) New() *Menu {
	// Get Central Time location (automatically handles daylight saving)
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		// Fallback to UTC if location loading fails
		loc = time.UTC
	}

	return &Menu{
		Items: []MenuItem{
			{Label: "Radar", URL: "/", Tooltip: "What you really came here for"},
			{Label: "Satellite", URL: "/satellite", Tooltip: ""},
			{Label: "Watches/Warnings", URL: "/watches", Tooltip: ""},
			{Label: "Current Temps", URL: "/temperatures", Tooltip: ""},
			{Label: "Outlook", URL: "/outlook", Tooltip: ""},
			{Label: "Rainfall Amounts", URL: "/rainfall", Tooltip: ""},
			{Label: "Resources", URL: "/resources", Tooltip: ""},
			{Label: "About", URL: "/about", Tooltip: ""},
			{Label: "Disclaimer", URL: "/disclaimer", Tooltip: ""},
			{Label: "Donate", URL: "/donate", Tooltip: "Buy me a coffee!"},
		},
		LoadTime: time.Now().In(loc),
	}
}

// IsSelected checks if a menu item is the current page
func (m *Menu) IsSelected(item MenuItem, currentPath string) bool {
	// Handle home page specially
	if item.URL == "/" && currentPath == "/" {
		return true
	}
	// For other pages, check if the current path starts with the item's URL
	return currentPath != "/" && filepath.Base(currentPath) == filepath.Base(item.URL)
}
