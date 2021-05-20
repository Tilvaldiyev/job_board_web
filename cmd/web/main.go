package main

import (
	"jobBoard/pkg/config"
	"jobBoard/pkg/handlers"
	"jobBoard/pkg/render"
	"log"
	"net/http"
)

const (
	portNumber = ":8080"
)

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// Set to global config created cache
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

