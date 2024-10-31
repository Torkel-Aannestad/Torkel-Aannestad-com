package main

import (
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"time"
)

type post struct {
	Title       string
	Slug        string
	Description string
	Date        string
}

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

	posts := map[string]post{}
	posts["post-design-system-something-test"] = post{
		Title:       "Crafting a design system for a multiplanetary future",
		Slug:        "post-design-system-something-test",
		Description: `Most companies try to stay ahead of the curve when it comes to visual design, but for Planetaria we needed to create a brand that would still inspire us 100 years from now when humanity has spread across our entire solar system.`,
		Date:        "September 5, 2022",
	}
	posts["tester-post"] = post{
		Title:       "TESTER POST",
		Slug:        "tester-post",
		Description: `Most companies try to stay ahead of the curve when it comes to visual design, but for Planetaria we needed to create a brand that would still inspire us 100 years from now when humanity has spread across our entire solar system.`,
		Date:        "Oktober 15, 2022",
	}

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
