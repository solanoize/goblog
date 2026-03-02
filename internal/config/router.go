package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	var r *chi.Mux = chi.NewRouter()

	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)

	return r
}
