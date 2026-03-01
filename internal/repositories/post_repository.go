package repositories

import (
	"context"

	"github.com/solanoize/goblog/internal/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAll(ctx context.Context, page int, limit int, search string) ([]models.Post, error)
	FindByID(ctx context.Context, id uint) (models.Post, error)
	Create(ctx context.Context, post models.Post) (models.Post, error)
	Update(ctx context.Context, post models.Post) (models.Post, error)
	Delete(ctx context.Context, id uint) error
}

type postRepository struct {
	DB                  *gorm.DB
	PaginatorRepository PaginatorRepository
}

// Create implements [PostRepository].
func (p *postRepository) Create(ctx context.Context, post models.Post) (models.Post, error) {
	var err error
	err = p.DB.Create(&post).Error

	if err != nil {
		return post, err
	}
	err = p.DB.Preload("User").First(&post, post.ID).Error
	return post, err
}

// Delete implements [PostRepository].
func (p *postRepository) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// FindAll implements [PostRepository].
func (p *postRepository) FindAll(ctx context.Context, page int, limit int, search string) ([]models.Post, error) {
	panic("unimplemented")
}

// FindByID implements [PostRepository].
func (p *postRepository) FindByID(ctx context.Context, id uint) (models.Post, error) {
	panic("unimplemented")
}

// Update implements [PostRepository].
func (p *postRepository) Update(ctx context.Context, post models.Post) (models.Post, error) {
	panic("unimplemented")
}

func NewPostRespository(db *gorm.DB) PostRepository {
	return &postRepository{
		DB:                  db,
		PaginatorRepository: NewPaginatorRepository(100),
	}
}
