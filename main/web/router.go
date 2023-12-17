package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/habib0071/goLang/internal/config"
	"github.com/habib0071/goLang/internal/handlers"

	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Simple)
	mux.Post("/post-simple", handlers.Repo.PostSimple)
	mux.Get("/simple-summary", handlers.Repo.SimpleSummary)


	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
