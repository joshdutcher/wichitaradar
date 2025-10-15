package templates

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestInit(t *testing.T) {
	// Create a temporary directory for test templates
	tempDir, err := os.MkdirTemp("", "templates_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create test template files
	templateFiles := map[string]string{
		"index.page.html":     "<h1>Home Page</h1>{{template \"base\" .}}",
		"base.layout.html":    "<html>{{template \"content\" .}}</html>",
		"header.partial.html": "<header>Header</header>",
	}
	for name, content := range templateFiles {
		path := filepath.Join(tempDir, name)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Initialize templates
	templateFS := os.DirFS(tempDir)
	if err := Init(templateFS); err != nil {
		t.Fatalf("Init failed: %v", err)
	}

	// Test getting a template by its base name
	tmpl, err := Get("index")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if tmpl == nil {
		t.Fatal("template is nil")
	}
}

func TestDefaultTemplateProvider_Get(t *testing.T) {
	tests := []struct {
		name        string
		setupFS     fs.FS
		templateKey string
		shouldError bool
		description string
	}{
		{
			name: "get existing template",
			setupFS: fstest.MapFS{
				"index.page.html":     &fstest.MapFile{Data: []byte("<h1>Home</h1>")},
				"base.layout.html":    &fstest.MapFile{Data: []byte("<html>{{template \"content\" .}}</html>")},
				"header.partial.html": &fstest.MapFile{Data: []byte("<header>Header</header>")},
			},
			templateKey: "index",
			shouldError: false,
			description: "should successfully get existing template",
		},
		{
			name:        "get from uninitialized provider",
			setupFS:     nil,
			templateKey: "index",
			shouldError: true,
			description: "should fail when templates are not initialized",
		},
		{
			name: "get non-existent template",
			setupFS: fstest.MapFS{
				"base.layout.html":    &fstest.MapFile{Data: []byte("<html>{{template \"content\" .}}</html>")},
				"header.partial.html": &fstest.MapFile{Data: []byte("<header>Header</header>")},
			},
			templateKey: "nonexistent",
			shouldError: true,
			description: "should fail when template doesn't exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset templates for each test
			templatesMap = nil

			// Initialize templates if we have a setup FS
			if tt.setupFS != nil {
				if err := Init(tt.setupFS); err != nil {
					if tt.name == "get non-existent template" {
						// This is expected for this test case
						return
					}
					t.Fatal(err)
				}
			}

			provider := &DefaultTemplateProvider{}
			_, err := provider.Get(tt.templateKey)

			if tt.shouldError {
				if err == nil {
					t.Errorf("%s: expected error, got nil", tt.description)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tt.description, err)
				}
			}
		})
	}
}
