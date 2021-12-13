package handler

import (
	"net/http"

	"github.com/mayurvpatil/go-server/pkg/config"
	"github.com/mayurvpatil/go-server/pkg/render"
)

// Repository variable
var Repo *Repository

// Repository is the reposirtory type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
