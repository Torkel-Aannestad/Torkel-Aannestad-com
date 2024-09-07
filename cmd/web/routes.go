package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(app.panicRecovery)
	router.Use(app.requestLogger)

	router.Get("/", app.home)

	return router
}
