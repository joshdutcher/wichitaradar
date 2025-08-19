package menu

import (
	"log"
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
	// This ensures Central Time is used regardless of server location
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Printf("Warning: Failed to load America/Chicago timezone: %v", err)
		// If timezone data is not available, try common alternatives
		loc, err = time.LoadLocation("US/Central")
		if err != nil {
			log.Printf("Warning: Failed to load US/Central timezone: %v", err)
			// Last resort: create fixed offset for Central Time
			// CST is UTC-6, but this won't handle daylight saving automatically
			loc = time.FixedZone("CST", -6*3600)
			log.Printf("Warning: Using fixed CST offset, daylight saving time will not be automatically handled")
		}
	}

	// Log the timezone being used for debugging
	now := time.Now().In(loc)
	log.Printf("Menu timestamp using timezone: %s, current time: %s", loc.String(), now.Format("2006-01-02 15:04:05 MST"))

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
		LoadTime: now,
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
