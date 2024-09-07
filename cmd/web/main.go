package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger slog.Logger
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := application{
		logger: *logger,
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
	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("server error", "error", err)
		os.Exit(1)
	}

}
