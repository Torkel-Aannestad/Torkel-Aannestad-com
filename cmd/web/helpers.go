package main

import "net/http"

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()
	app.logger.Error("server error", "method", method, "uri", uri, "error", err)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientErrorResponse(w http.ResponseWriter, statusCode int) {
	if statusCode < 400 || statusCode > 499 {
		statusCode = 400
	}

	http.Error(w, http.StatusText(statusCode), statusCode)
}
