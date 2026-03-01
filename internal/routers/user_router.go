package routers

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/solanoize/goblog/internal/controllers"
	"github.com/solanoize/goblog/internal/middlewares"
	"github.com/solanoize/goblog/internal/repositories"
	"github.com/solanoize/goblog/internal/usecases"
	"gorm.io/gorm"
)

type UserRouter interface {
	Register() *chi.Mux
}

type userRouter struct {
	DB     *gorm.DB
	Logger *log.Logger
	Router *chi.Mux
}

// Register implements [UserRouter].
func (u *userRouter) Register() *chi.Mux {
	var userRepository repositories.UserRepository = repositories.NewUserRepository(u.DB)
	var userUseCase usecases.UserUseCase = usecases.NewUserUseCase(u.Logger, userRepository)
	var userController controllers.UserController = controllers.NewUserController(userUseCase)
	var authMiddleware middlewares.AuthMiddleware = middlewares.NewAuthMiddleware(u.Logger, userUseCase)

	u.Router.With(
		authMiddleware.IsAuthenticated(),
		authMiddleware.IsActiveUser(),
	).Get("/me", userController.Me)

	u.Router.With(
		authMiddleware.IsAuthenticated(),
		authMiddleware.IsActiveUser(),
		authMiddleware.IsAdminOnly(),
	).Get("/users", userController.List)

	u.Router.Post("/register", userController.Register)
	u.Router.Post("/signin", userController.SignIn)

	return u.Router
}

func NewUserRouter(db *gorm.DB, logger *log.Logger) UserRouter {
	return &userRouter{
		DB:     db,
		Logger: logger,
		Router: chi.NewRouter(),
	}
}
