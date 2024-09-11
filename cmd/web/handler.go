package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl", "home", nil)
}
func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "about.tmpl", "about", nil)
}
func (app *application) postHandler(w http.ResponseWriter, r *http.Request) {
	data := []string{"1", "2", "3"}
	app.render(w, r, http.StatusOK, "posts.tmpl", "posts", data)
}
func (app *application) postDetailsHandler(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, http.StatusOK, "post-details-slug.tmpl", "post-details", nil)
}
