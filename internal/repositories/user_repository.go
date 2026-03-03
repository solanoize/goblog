package repositories

import (
	"context"

	"github.com/solanoize/goblog/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

// Create implements [UserRepository].
func (u *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := u.DB.WithContext(ctx).Create(&user).Error
	return user, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func UserRegisterRepository(db *gorm.DB, ctx context.Context, user models.User) {

}
