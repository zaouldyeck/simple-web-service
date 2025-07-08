package handlers

import (
	"net/http"

	"github.com/zaouldyeck/simple-web-service/pkg/render"
)

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.tmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.tmpl")
}
