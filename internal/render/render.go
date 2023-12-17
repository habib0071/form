package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/habib0071/goLang/internal/config"
	"github.com/habib0071/goLang/internal/model"
)

var function = template.FuncMap{}

var app *config.AppConfig

//NewTemplates sets the confiq for the template package

func Newtemplates(a *config.AppConfig) {
	app = a
}

func AddDefaltData(td *model.TemplateData) *model.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCashe {
		// get the templates cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaltData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser, err")
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// myCache := make(map[string]*Template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(("./templates/*.layout.tmpl"))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, err
}
