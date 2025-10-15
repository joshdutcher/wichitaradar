package menu

import (
	"testing"
)

func TestNew(t *testing.T) {
	menu := New()

	// Check that all expected menu items are present
	expectedItems := []string{
		"Radar",
		"Satellite",
		"Watches/Warnings",
		"Current Temps",
		"Outlook",
		"Rainfall Amounts",
		"Resources",
		"About",
		"Disclaimer",
		"Donate",
	}

	if len(menu.Items) != len(expectedItems) {
		t.Errorf("expected %d menu items, got %d", len(expectedItems), len(menu.Items))
	}

	for i, item := range menu.Items {
		if item.Label != expectedItems[i] {
			t.Errorf("expected menu item %q, got %q", expectedItems[i], item.Label)
		}
	}

	// Check that LoadTime is set
	if menu.LoadTime.IsZero() {
		t.Error("expected LoadTime to be set")
	}
}

func TestIsSelected(t *testing.T) {
	menu := New()

	tests := []struct {
		name        string
		item        MenuItem
		currentPath string
		want        bool
	}{
		{
			name:        "home page selected",
			item:        MenuItem{URL: "/"},
			currentPath: "/",
			want:        true,
		},
		{
			name:        "home page not selected",
			item:        MenuItem{URL: "/"},
			currentPath: "/about",
			want:        false,
		},
		{
			name:        "about page selected",
			item:        MenuItem{URL: "/about"},
			currentPath: "/about",
			want:        true,
		},
		{
			name:        "about page not selected",
			item:        MenuItem{URL: "/about"},
			currentPath: "/",
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := menu.IsSelected(tt.item, tt.currentPath); got != tt.want {
				t.Errorf("IsSelected() = %v, want %v", got, tt.want)
			}
		})
	}
}
