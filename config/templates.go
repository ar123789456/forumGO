package config

import "html/template"

var Tmpl *template.Template

func ParseTemplates() (*template.Template, error) {
	var err error
	Tmpl, err = template.ParseGlob("templates/*.html")
	return Tmpl, err
}
