package users

import (
	"log"

	"github.com/solanoize/goblog/internal/auth"
	"gorm.io/gorm"
)

type Resource interface {
	Apply()
	GetModel() User
	GetRepository() Repository
	GetService() Service
	GetController() Controller
	GetMiddleware() Middleware
}

type resource struct {
	Model      User
	Repository Repository
	Service    Service
	Controller Controller
	Middleware Middleware

	Infrastructure struct {
		DB     *gorm.DB
		Logger *log.Logger
	}

	Dependency struct {
		Auth auth.Resource
	}
}

// GetController implements [Resource].
func (r *resource) GetController() Controller {
	return r.Controller
}

// GetMiddleware implements [Resource].
func (r *resource) GetMiddleware() Middleware {
	return r.Middleware
}

// GetModel implements [Resource].
func (r *resource) GetModel() User {
	return r.Model
}

// GetRepository implements [Resource].
func (r *resource) GetRepository() Repository {
	return r.Repository
}

// GetService implements [Resource].
func (r *resource) GetService() Service {
	return r.Service
}

// Apply implements [Resource].
func (r *resource) Apply() {
	model := User{}
	repository := NewRepository(r.Infrastructure.DB)
	service := NewService(repository, r.Dependency.Auth.GetService(), r.Infrastructure.Logger)
	controller := NewController(service)
	middleware := NewMiddleware(service, r.Infrastructure.Logger)

	r.Model = model
	r.Repository = repository
	r.Service = service
	r.Controller = controller
	r.Middleware = middleware
}

func NewResource(db *gorm.DB, logger *log.Logger, authResource auth.Resource) Resource {
	return &resource{
		Infrastructure: struct {
			DB     *gorm.DB
			Logger *log.Logger
		}{
			DB:     db,
			Logger: logger,
		},
		Dependency: struct{ Auth auth.Resource }{
			Auth: authResource,
		},
	}
}
