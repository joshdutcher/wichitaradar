package templates

import (
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

// Store templates in a map, keyed by the page template filename (e.g., "index.page.html")
var templatesMap map[string]*template.Template

// DefaultTemplateProvider is the default implementation of TemplateProvider
type DefaultTemplateProvider struct{}

// Get returns the template set for a specific page
func (p *DefaultTemplateProvider) Get(pageName string) (*template.Template, error) {
	// Ensure .page.html suffix if not present
	if !strings.HasSuffix(pageName, ".page.html") {
		pageName += ".page.html"
	}
	ts, ok := templatesMap[pageName]
	if !ok {
		return nil, fmt.Errorf("template %q not found", pageName)
	}
	return ts, nil
}

func Init(templateFS fs.FS) error {
	templatesMap = make(map[string]*template.Template)

	// 1. Get all page template paths
	pages, err := fs.Glob(templateFS, "*.page.html")
	if err != nil {
		return fmt.Errorf("failed to glob page templates: %w", err)
	}
	if len(pages) == 0 {
		return fmt.Errorf("no *.page.html files found in templates directory")
	}

	// 2. For each page template, parse it along with layouts and partials
	for _, page := range pages {
		name := filepath.Base(page)

		// Create a new template set for this page
		ts, err := template.New(name).ParseFS(templateFS,
			page,             // The specific page template
			"*.layout.html",  // All base layouts
			"*.partial.html", // All partials/components
		)
		if err != nil {
			return fmt.Errorf("failed to parse template set for %s: %w", name, err)
		}

		templatesMap[name] = ts
	}

	fmt.Printf("Initialized %d page templates\n", len(templatesMap))
	return nil
}

// Get returns the template set for a specific page using the default provider
func Get(pageName string) (*template.Template, error) {
	return (&DefaultTemplateProvider{}).Get(pageName)
}
