package main

import "net/http"

func (app *application) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		uri := r.URL.RequestURI()
		ip := r.RemoteAddr
		app.logger.Info("request received", "uri", uri, "method", method, "ip", ip)

		next.ServeHTTP(w, r)
	})
}
