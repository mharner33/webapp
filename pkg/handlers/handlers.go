package handlers

import (
	"net/http"

	"github.com/mharner33/webapp/pkg/config"
	"github.com/mharner33/webapp/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Create a new repository to easily swap out config options
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
