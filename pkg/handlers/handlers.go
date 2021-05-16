package handlers

import (
	"jobBoard/pkg/render"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {

}