package main

import (
	"net/http"

	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/habib0071/goLang/internal/config"
	"github.com/habib0071/goLang/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSruve)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Simple)
	mux.Get("/simple-form", handlers.Repo.PostSimple)
	//mux.Get("/about", handlers.Repo.About)
	// mux.Get("/contact", handlers.Repo.Contact)
	// mux.Post("/contact-form", handlers.Repo.PostContact)

	// mux.Get("/fashion", handlers.Repo.Fashion)
	// mux.Get("/travel", handlers.Repo.Travel)
	// mux.Get("/single", handlers.Repo.Single)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
