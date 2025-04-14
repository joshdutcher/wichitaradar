package handlers

import (
	"testing"

	"wichitaradar/internal/testutils"
)

func TestHandleSimplePage(t *testing.T) {
	tests := []struct {
		name         string
		templateName string
		path         string
	}{
		{
			name:         "about page",
			templateName: "about",
			path:         "/about",
		},
		{
			name:         "disclaimer page",
			templateName: "disclaimer",
			path:         "/disclaimer",
		},
		{
			name:         "donate page",
			templateName: "donate",
			path:         "/donate",
		},
		{
			name:         "resources page",
			templateName: "resources",
			path:         "/resources",
		},
		{
			name:         "watches page",
			templateName: "watches",
			path:         "/watches",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.InitTemplates(t)
			testutils.TestHandler(t, HandleSimplePage(tt.templateName), tt.path)
		})
	}
}