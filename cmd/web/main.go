package main

import (
	"bulletinBoard/pkg/config"
	"bulletinBoard/pkg/handlers"
	"bulletinBoard/pkg/render"
	"log"
	"net/http"
)

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.HomePage)

	http.ListenAndServe(":8080", nil)

}

