package handlers

import (
	"net/http"

	"github.com/gabrielyea/mini-booking-go/pkg/config"
	"github.com/gabrielyea/mini-booking-go/pkg/models"
	"github.com/gabrielyea/mini-booking-go/pkg/render"
)

var Repo *Repository

// Repository is the reposotory type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	re.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	sm := make(map[string]string)
	sm["test"] = "Hello there home"
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{StringMap: sm})
}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := re.App.Session.GetString(r.Context(), "remote_ip")

	sm := make(map[string]string)
	sm["remote_ip"] = remoteIp
	sm["test"] = "Hello there about"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: sm})
}
