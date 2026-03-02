package config

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Server(env string, port string, router *chi.Mux, logger *log.Logger) {
	logger.Printf("Server running on environment [%s] at port %s\n", env, port)

	var err error = http.ListenAndServe(":"+port, router)
	if err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}
