package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/habib0071/goLang/internal/config"
	"github.com/habib0071/goLang/internal/handlers"
	"github.com/habib0071/goLang/internal/render"
)

const protNumber = ":8012"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	var app config.AppConfig

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
    session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCashe = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.Newtemplates(&app)

	fmt.Println(fmt.Sprintf("The application starting prot number is %s", protNumber))

    srv := &http.Server {
		Addr: protNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
