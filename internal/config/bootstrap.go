package config

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/solanoize/goblog/internal/controllers"
	"github.com/solanoize/goblog/internal/repositories"
	"github.com/solanoize/goblog/internal/services"
	"gorm.io/gorm"
)

func Bootstrap(db *gorm.DB, logger *log.Logger, router chi.Router) {

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(logger, userRepo)
	userController := controllers.NewUserController(userService)

	router.Route("/", func(r chi.Router) {
		r.Post("/register", userController.SignUp)
	})

}
