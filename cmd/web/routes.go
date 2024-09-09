package main

import (
	"net/http"

	"github.com/Torkel-Aannestad/torkel-dev/ui"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(app.requestLogger)
	r.Use(app.panicRecovery)
	r.Use(app.commonHeaders)

	r.Method(http.MethodGet, "/public/*", http.FileServerFS(ui.Files))

	r.Get("/", app.homeHandler)
	r.Get("/about", app.aboutHandler)
	return r
}
