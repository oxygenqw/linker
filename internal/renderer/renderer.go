package renderer

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	files, err := filepath.Glob("./web/pages/*.html")
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)
	for _, file := range files {
		name := filepath.Base(file)
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			return nil, err
		}
		templates[name] = tmpl
	}

	return &TemplateRenderer{templates: templates}, nil
}

func (r *TemplateRenderer) RenderTemplate(w http.ResponseWriter, tmplName string, data any) {
	tmpl, ok := r.templates[tmplName]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
