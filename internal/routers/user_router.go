package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/solanoize/goblog/internal/controllers"
	"github.com/solanoize/goblog/internal/middlewares"
)

type UserRouter interface {
	Register()
}

type userRouter struct {
	Router         chi.Router
	AuthMidleware  middlewares.AuthMiddleware
	UserController controllers.UserController
}

// Register implements [UserRouter].
func (u *userRouter) Register() {

	u.Router.Route("/", func(r chi.Router) {
		r.Post("/register", u.UserController.Register)
		r.Post("/signin", u.UserController.SignIn)

		r.Group(func(r chi.Router) {
			r.Use(u.AuthMidleware.IsAuthenticated())
			r.Use(u.AuthMidleware.IsActiveUser())

			r.With(u.AuthMidleware.IsAdminOnly()).Get("/users", u.UserController.List)
			r.Get("/users/me", u.UserController.Me)
		})
	})
}

func NewUserRouter(router chi.Router, authMiddleware middlewares.AuthMiddleware, userController controllers.UserController) UserRouter {
	return &userRouter{
		Router:         router,
		AuthMidleware:  authMiddleware,
		UserController: userController,
	}
}
