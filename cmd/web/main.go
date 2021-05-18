package main

import (
	"jobBoard/pkg/config"
	"jobBoard/pkg/handlers"
	"jobBoard/pkg/render"
	"log"
	"net/http"
)

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// Set to global config created cache
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.HomePage)

	http.ListenAndServe(":8080", nil)

}

