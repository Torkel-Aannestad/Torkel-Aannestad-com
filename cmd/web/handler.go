package main

import (
	"fmt"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl", nil)
}
func (app *application) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "not-found.tmpl", nil)
}
func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "about.tmpl", nil)
}
func (app *application) postHandler(w http.ResponseWriter, r *http.Request) {
	data := []post{}
	for _, v := range app.postsData {
		data = append(data, v)
	}

	app.render(w, r, http.StatusOK, "posts.tmpl", data)
}

func (app *application) postDetailsHandler(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("id")
	_, ok := app.postsData[path]
	if !ok {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
		return
	}

	pathFile := fmt.Sprintf("%v.tmpl", path)
	app.render(w, r, http.StatusOK, pathFile, nil)
}
