package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/solanoize/goblog/internal/controllers"
	"github.com/solanoize/goblog/internal/middlewares"
)

type PostRouter interface {
	Register()
}

type postRouter struct {
	Router         chi.Router
	AuthMiddleware middlewares.AuthMiddleware
	PostController controllers.PostController
}

// Register implements [PostRouter].
func (p *postRouter) Register() {
	p.Router.Route("/posts", func(r chi.Router) {
		r.Use(p.AuthMiddleware.IsAuthenticated())
		r.Use(p.AuthMiddleware.IsActiveUser())
		r.Post("/", p.PostController.Create)
	})
}

func NewPostRouter(router chi.Router, authMiddleware middlewares.AuthMiddleware, postController controllers.PostController) PostRouter {
	return &postRouter{
		Router:         router,
		AuthMiddleware: authMiddleware,
		PostController: postController,
	}
}
