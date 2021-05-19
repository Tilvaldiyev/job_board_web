package config

import "html/template"

// AppConfig holds application configs
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
