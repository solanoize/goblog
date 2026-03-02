package auth

import (
	"log"

	"gorm.io/gorm"
)

type Resource interface {
	Apply()
	GetService() Service
}

type resource struct {
	Service        Service
	Infrastructure struct {
		DB     *gorm.DB
		Logger *log.Logger
	}
}

// GetService implements [Resource].
func (r *resource) GetService() Service {
	return r.Service
}

// Apply implements [Resource].
func (r *resource) Apply() {
	r.Service = NewService(r.Infrastructure.Logger)
}

func NewResource(db *gorm.DB, logger *log.Logger) Resource {
	return &resource{
		Infrastructure: struct {
			DB     *gorm.DB
			Logger *log.Logger
		}{
			DB:     db,
			Logger: logger,
		},
	}
}
