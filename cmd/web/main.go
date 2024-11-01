package main

import (
	"expvar"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"text/template"
	"time"
)

type config struct {
	port int
	env  string
	// rateLimiter struct {
	// 	rps     float64
	// 	burst   int
	// 	enabled bool
	// }
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
	//rate limiter
	// flag.Float64Var(&cfg.rateLimiter.rps, "rate-limit-rps", 2, "request per minute per ip")
	// flag.IntVar(&cfg.rateLimiter.burst, "rate-limit-burst", 4, "rate limit bucket size")
	// flag.BoolVar(&cfg.rateLimiter.enabled, "rate-limit-enabled", true, "rate limiter toggle")

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

	//expvar metrics endpoint
	// expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))
	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

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
