package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"jobBoard/pkg/config"
	"jobBoard/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)

	return mux
}
