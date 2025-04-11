package templates

import "html/template"

var Templates *template.Template

func init() {
	// Load all templates
	Templates = template.Must(template.ParseGlob("templates/*.html"))
}
