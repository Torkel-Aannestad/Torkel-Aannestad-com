package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl", "home", nil)
}
func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "about.tmpl", "about", nil)
}
func (app *application) articleHandler(w http.ResponseWriter, r *http.Request) {
	data := []string{"1", "2", "3"}
	app.render(w, r, http.StatusOK, "articles.tmpl", "articles", data)
}
func (app *application) articleDetailsHandler(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, http.StatusOK, "article-details-slug.tmpl", "article-details", nil)
}
