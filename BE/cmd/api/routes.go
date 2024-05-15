package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// creating a route mux
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// routes
	mux.Get("/authenticate", app.authenticate) // mocked user
	mux.Get("/", app.Home)
	mux.Get("/movies", app.getAllMovies)

	return mux
}
