package templates

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	// Create a temporary directory for test templates
	tempDir, err := os.MkdirTemp("", "templates_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create test template files
	templateFiles := []struct {
		name     string
		content  string
	}{
		{"index.page.html", "<h1>Home Page</h1>"},
		{"base.layout.html", "<html>{{template \"content\" .}}</html>"},
		{"header.partial.html", "<header>Header</header>"},
	}

	for _, tf := range templateFiles {
		path := filepath.Join(tempDir, tf.name)
		if err := os.WriteFile(path, []byte(tf.content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Initialize templates
	templateFS := os.DirFS(tempDir)
	if err := Init(templateFS); err != nil {
		t.Fatal(err)
	}

	// Test Get with existing template
	tmpl, err := Get("index")
	if err != nil {
		t.Fatal(err)
	}
	if tmpl == nil {
		t.Error("expected non-nil template")
	}

	// Test Get with non-existent template
	_, err = Get("nonexistent")
	if err == nil {
		t.Error("expected error for non-existent template")
	}
}

func TestDefaultTemplateProvider_Get(t *testing.T) {
	// Reset templatesMap to nil to ensure we're testing uninitialized state
	templatesMap = nil

	provider := &DefaultTemplateProvider{}

	// Test with .page.html suffix
	_, err := provider.Get("index.page.html")
	if err == nil {
		t.Error("expected error when templates not initialized")
	}

	// Test without .page.html suffix
	_, err = provider.Get("index")
	if err == nil {
		t.Error("expected error when templates not initialized")
	}
}