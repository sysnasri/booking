package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sysnasri/booking/pkg/config"
	"github.com/sysnasri/booking/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(SessionLoad)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)
	mux.Use(WriteToConsole)
	mux.Use(LogIPAddress)
	mux.Use(NoSrve)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	return mux

}
