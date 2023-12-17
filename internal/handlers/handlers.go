package handlers

import (
	"log"
	"net/http"

	"github.com/habib0071/goLang/internal/config"
	"github.com/habib0071/goLang/internal/forms"
	"github.com/habib0071/goLang/internal/model"
	"github.com/habib0071/goLang/internal/render"
)

// Repo the reository used by the handler
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Simple(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "simple.page.tmpl", &model.TemplateData{})
}

func (m *Repository) PostSimple(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return 
	}

	simpleInformation := model.SimpleInformation{
		UserName: r.Form.Get("username"),
		Email: r.Form.Get("email"),
		Password: r.Form.Get("password"),
		Cpassword: r.Form.Get("cpassword"),
	}
	form := forms.New(r.PostForm)

	form.Has("full_name", r)

	if form.Valid() {
		data := make(map[string]interface{})
		data["simpleInformation"] = simpleInformation

		render.RenderTemplate(w, "simple.page.tmpl", &model.TemplateData{
			Form: form,
			Data: data,

		})
		return
	}
}


