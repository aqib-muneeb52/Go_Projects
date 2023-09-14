// templates/template_loader.go

package templates

import (
	"html/template"
)

// LoadTemplates loads your templates
func LoadTemplates() (*template.Template, error) {
	// Load and parse your templates here
	return template.ParseGlob("templates/*.html")
}
