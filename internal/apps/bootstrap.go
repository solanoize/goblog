package apps

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/solanoize/goblog/internal/auth"
	"github.com/solanoize/goblog/internal/users"
	"github.com/solanoize/goblog/internal/utils"
	"gorm.io/gorm"
)

type Bootstrap interface {
	Wire()
	Routing()
	Migrate()
}

type bootstrap struct {
	DB        *gorm.DB
	Logger    *log.Logger
	Router    *chi.Mux
	Resources struct {
		Auth auth.Resource
		User users.Resource
	}
}

// Migrate implements [Bootstrap].
func (b *bootstrap) Migrate() {
	models := []interface{}{
		b.Resources.User.GetModel(),
	}

	if err := b.DB.AutoMigrate(models...); err != nil {
		b.Logger.Fatal("Gagal migrasi database: ", err)
	}

	b.Logger.Println("Database migration completed successfully")
}

// Routing implements [Bootstrap].
func (b *bootstrap) Routing() {
	b.Router.Route("/", func(r chi.Router) {
		r.Post("/register", b.Resources.User.GetController().Create)
		r.Post("/signin", b.Resources.User.GetController().SignIn)
		r.Group(func(r chi.Router) {
			r.Use(utils.MiddlewareAdapter(b.Resources.User.GetMiddleware().IsAuthenticated))
			r.Get("/users/me", b.Resources.User.GetController().Retrieve)
		})
	})
}

// Register implements [Bootstrap].
func (b *bootstrap) Wire() {
	authResource := auth.NewResource(b.DB, b.Logger)
	userResource := users.NewResource(b.DB, b.Logger, authResource)

	b.Resources.Auth = authResource
	b.Resources.User = userResource

	b.Resources.Auth.Apply()
	b.Resources.User.Apply()
}

func NewBootstrap(db *gorm.DB, logger *log.Logger, router *chi.Mux) Bootstrap {
	return &bootstrap{
		DB:     db,
		Logger: logger,
		Router: router,
	}
}
