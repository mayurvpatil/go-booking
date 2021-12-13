package config

import "text/template"

// AppConfig holds application configurations
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
