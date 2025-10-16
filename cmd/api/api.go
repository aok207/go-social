package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}
type config struct {
	addr string
}

func (app *application) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	r.Use(middleware.Timeout(time.Second * 60))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/_/sys/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) serve(handler http.Handler) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
	}

	log.Printf("Server is running at port%s\n", app.config.addr)
	return server.ListenAndServe()
}
