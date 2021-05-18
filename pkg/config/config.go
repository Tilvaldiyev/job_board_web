package config

import "html/template"

// AppConfig holds application configs
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
