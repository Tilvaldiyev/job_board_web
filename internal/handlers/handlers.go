package handlers

import (
	"fmt"
	"jobBoard/internal/config"
	"jobBoard/internal/forms"
	"jobBoard/internal/models"
	"jobBoard/internal/render"
	"log"
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
	var emptyApplyJob models.ApplyJob
	data := make(map[string]interface{})
	data["applyJob"] = emptyApplyJob

	render.RenderTemplate(w, r, "job_details.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// ApplyJob is the job description page handler
func (repo *Repository) ApplyJob(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println(err)
		return
	}

	//file, handler, err := r.FormFile("cv_file")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()

	//f, err := os.OpenFile("/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()

	applyJob := models.ApplyJob{
		FirstName: r.Form.Get("name"),
		Email: r.Form.Get("email"),
		WebsiteLink: r.Form.Get("website"),
		Coverletter: r.Form.Get("coverletter"),
		//Portfolio: f,
	}

	form := forms.New(r.PostForm)

	form.Required("name", "email", "website")
	form.MinLength("name", 3, r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["applyJob"] = applyJob

		render.RenderTemplate(w, r, "job_details.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	//io.Copy(f, file)
	fmt.Println(applyJob.FirstName)

}

// SingleBlogPage is the single blog page handler
func (repo *Repository) SingleBlogPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "single_blog.page.tmpl", &models.TemplateData{})
}
