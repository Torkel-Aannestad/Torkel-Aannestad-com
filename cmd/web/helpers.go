package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()
	app.logger.Error("server error", "method", method, "uri", uri, "error", err)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// func (app *application) clientErrorResponse(w http.ResponseWriter, statusCode int) {
// 	if statusCode < 400 || statusCode > 499 {
// 		statusCode = 400
// 	}

// 	http.Error(w, http.StatusText(statusCode), statusCode)
// }

func (app *application) render(w http.ResponseWriter, r *http.Request, statusCode int, page string, data interface{}) {
	templateSet, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("could not find requested page template in render")
		app.serverErrorResponse(w, r, err)
		return
	}

	// if r.Header.Get("HX-Request") == "true" {
	// 	w.Header().Set("Vary", "HX-Request")
	// }
	w.Header().Set("Cache-Control", "public")

	template := strings.Split(page, ".")
	buf := new(bytes.Buffer)
	err := templateSet.ExecuteTemplate(buf, template[0], data)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(statusCode)

	buf.WriteTo(w)

}
