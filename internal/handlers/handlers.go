package handlers

import (
	"log"
	"net/http"

	"github.com/habib0071/goLang/internal/config"
	"github.com/habib0071/goLang/internal/forms"
	"github.com/habib0071/goLang/internal/models"
	"github.com/habib0071/goLang/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
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
func (m *Repository) Simple(w http.ResponseWriter, r *http.Request) {
	var emtySimpleform models.Simpleform     

	data := make(map[string]interface{})

	data["simpleform"] = emtySimpleform
	render.RenderTemplate(w, r, "simple.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}	

func (m *Repository) PostSimple(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}	
	simpleform := models.Simpleform {
		Username:   r.Form.Get("username"),
		Email:      r.Form.Get("email"),
	}
	form := forms.New(r.PostForm)

	form.Required("username", "email",)
	form.MinLength("username", 3, r)
	form.IsEmail("email")
    

	if !form.Valid() {
		data := make(map[string]interface{})
		data["simpleform"] = simpleform

		render.RenderTemplate(w, r, "simple.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	} 
	
	m.App.Session.Put(r.Context(), "simpleform", simpleform)
	http.Redirect(w, r, "/simple-summary", http.StatusSeeOther)
}

func (m *Repository) SimpleSummary(w http.ResponseWriter, r *http.Request) {
	simpleform, ok := m.App.Session.Get(r.Context(), "simpleform").(models.Simpleform)

	if !ok {
		log.Println("Cannot get item from Session")
		m.App.Session.Put(r.Context(), "error", "Can't get simpleform from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} 
	m.App.Session.Remove(r.Context(), "simpleform")
  	data := make(map[string]interface{})   
	data["simpleform"] = simpleform
	render.RenderTemplate(w, r, "post-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

