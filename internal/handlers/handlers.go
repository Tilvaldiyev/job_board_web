package handlers

import (
	"fmt"
	"io"
	"jobBoard/internal/config"
	"jobBoard/internal/forms"
	"jobBoard/internal/models"
	"jobBoard/internal/render"
	"net/http"
	"os"
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
	render.RenderTemplate(w, r, "job_details.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// ApplyJob is the job description page handler
func (repo *Repository) ApplyJob(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	website := r.Form.Get("website")
	coverletter := r.Form.Get("coverletter")

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("cv_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Println(name, email, website, coverletter)
}

// SingleBlogPage is the single blog page handler
func (repo *Repository) SingleBlogPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "single_blog.page.tmpl", &models.TemplateData{})
}
