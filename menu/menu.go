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
	var loc *time.Location
	var now time.Time
	
	// Try to load timezone data, but don't spam logs on every menu creation
	if tzLoc, err := time.LoadLocation("America/Chicago"); err == nil {
		loc = tzLoc
	} else if tzLoc, err := time.LoadLocation("US/Central"); err == nil {
		loc = tzLoc
	} else {
		// Last resort: create fixed offset for Central Time
		// Since we can't access timezone rules, always use CDT (-5) during deployment
		// This is more accurate than CST during the daylight saving period (Mar-Nov)
		loc = time.FixedZone("CDT", -5*3600) // Central Daylight Time
	}
	
	now = time.Now().In(loc)

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
