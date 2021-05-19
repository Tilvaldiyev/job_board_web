package render

import (
	"bytes"
	"html/template"
	"jobBoard/pkg/config"
	"log"
	"net/http"
	"path/filepath"
)

// Map of functions that we can use in templates
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl] // Get cached template
	if !ok {
		log.Fatal("Could not get tempate")
	}

	buf := new(bytes.Buffer)

	// Save it in buffer
	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w) // Get response
	if err != nil {
		w.Write([]byte("Error while writing template to browser: " + err.Error()))
	}
}

// Creates a templates cahce as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// Find all files with starting with anything and end with page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // Get only file name (page is an absolute path to the file)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page) // Create new template
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl") // Layouts that's defined for this template
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // Parse glob layouts for this template set
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
