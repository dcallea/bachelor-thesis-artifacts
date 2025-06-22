package handlers

import (
	"embed"
	"html/template"
	"net/http"
	//"path/filepath"
)

// references the embedded FS
var fs embed.FS

// set our local fs variable to point to the embedded filesystem
func SetFS(f embed.FS) {
	fs = f
}

// helper to load templates
func loadTemplate(tmplName string) (*template.Template, error) {
	//tmplPath := filepath.Join("templates", tmplName)
	tmpl, err := template.ParseFS(fs, "templates/"+tmplName)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// helper to render templates
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := loadTemplate(tmplName)
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render the page", http.StatusInternalServerError)
	}
}
