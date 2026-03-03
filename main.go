package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/solanoize/goblog/internal/config"
	"github.com/solanoize/goblog/internal/globals"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" || appEnv == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Lagi dev tapi file .env gak ketemu, cek duls!")
		}
	}

	// db := config.Postgre()
	// logger := config.Logging()
	// mainRouter := config.Router()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5000"
	}
	if appEnv == "" {
		appEnv = "development (default)"
	}

	// bootstrap := apps.NewBootstrap(db, logger, mainRouter)

	// bootstrap.Wire()
	// bootstrap.Migrate()
	// bootstrap.Routing()

	// config.Server(appEnv, port, mainRouter, logger)

	db := config.Postgre()
	globals.GlobalDB = db
	logger := config.Logging()
	router := config.Router()

	config.Bootstrap(db, logger, router)
	config.Server(appEnv, port, router, logger)
}
