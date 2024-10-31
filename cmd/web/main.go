package main

import (
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"time"
)

type application struct {
	logger        slog.Logger
	templateCache map[string]*template.Template
	postsData     map[string]post
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error("unable to get templates", "error", err)
		os.Exit(1)
	}

	posts := newPostsData()

	app := application{
		logger:        *logger,
		templateCache: templateCache,
		postsData:     posts,
	}

	srv := http.Server{
		Addr:         ":4000",
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting server", "addr", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("server error", "error", err)
		os.Exit(1)
	}

}
