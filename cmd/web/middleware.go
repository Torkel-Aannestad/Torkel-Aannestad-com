package main

import (
	"fmt"
	"mime"
	"net/http"
	"path"
)

func (app *application) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		uri := r.URL.RequestURI()
		ip := r.RemoteAddr
		app.logger.Info("request received", "uri", uri, "method", method, "ip", ip)

		next.ServeHTTP(w, r)
	})
}

func (app *application) panicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				err = fmt.Errorf("%v", err)
				app.logger.Error("panic recovery", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("hot dang, this failed somehow"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; "+
				"script-src 'self' https://analytics.easywave.io https://unpkg.com https://cdn.jsdelivr.net; "+
				"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com/ http://localhost:4000; "+
				"font-src https://fonts.gstatic.com;")

		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		ext := path.Ext(r.URL.Path)
		if ext != "" {
			contentType := mime.TypeByExtension(ext)
			if contentType != "" {
				w.Header().Set("Content-Type", contentType)
			}
		} else {
			w.Header().Set("Content-Type", "text/html")
		}
		// w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}
