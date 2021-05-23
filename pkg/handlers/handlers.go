package handlers

import (
	"fmt"
	"jobBoard/pkg/config"
	"jobBoard/pkg/models"
	"jobBoard/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repsotory
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"

	render.RenderTemplate(w, r, "index.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// JobsPage is the jobs list handler
func (repo *Repository) JobsPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "jobs.page.tmpl", &models.TemplateData{})
}

// CandidatesPage is the candidates list handler
func (repo *Repository) CandidatesPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "candidate.page.tmpl", &models.TemplateData{})
}

// BlogPage is the blogs list handler
func (repo *Repository) BlogPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "blog.page.tmpl", &models.TemplateData{})
}

// ContactPage is the contact page handler
func (repo *Repository) ContactPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// JobDetailsPage is the job description page handler
func (repo *Repository) JobDetailsPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "job_details.page.tmpl", &models.TemplateData{})
}

// ApplyJob is the job description page handler
func (repo *Repository) ApplyJob(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	website := r.Form.Get("website")
	file_cv := r.Form.Get("file_cv")
	coverletter := r.Form.Get("coverletter")
	fmt.Println(name, email, website, file_cv, coverletter)
	//w.Write([]byte("Posted"))
}

// SingleBlogPage is the single blog page handler
func (repo *Repository) SingleBlogPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "single_blog.page.tmpl", &models.TemplateData{})
}
