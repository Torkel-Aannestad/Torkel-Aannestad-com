package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config        config
	logger        slog.Logger
	templateCache map[string]*template.Template
	postsData     map[string]post
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	//cors
	// flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
	// 	cfg.cors.trustedOrigins = strings.Fields(val)
	// 	return nil
	// })

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error("unable to get templates", "error", err)
		os.Exit(1)
	}

	posts := newPostsData()

	app := application{
		config:        cfg,
		logger:        *logger,
		templateCache: templateCache,
		postsData:     posts,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", app.config.port),
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
