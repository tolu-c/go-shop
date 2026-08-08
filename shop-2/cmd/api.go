package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	config config
}

type config struct {
	port string
	db dbConfig
}

type dbConfig struct {
	dbUrl string
}

func (a *app) mount () http.Handler {
	r := chi.NewRouter()
	
	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	// http.ListenAndServe(":3333", r)
}