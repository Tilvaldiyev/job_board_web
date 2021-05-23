package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"jobBoard/internal/config"
	"jobBoard/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/jobs", handlers.Repo.JobsPage)
	mux.Get("/candidates", handlers.Repo.CandidatesPage)
	mux.Get("/blog", handlers.Repo.BlogPage)
	mux.Get("/contact", handlers.Repo.ContactPage)
	mux.Get("/job_details", handlers.Repo.JobDetailsPage)
	mux.Post("/job_details", handlers.Repo.ApplyJob)
	mux.Get("/single_blog", handlers.Repo.SingleBlogPage)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
