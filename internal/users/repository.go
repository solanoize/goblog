package users

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id uint) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id uint) error
}

type repository struct {
	DB *gorm.DB
}

// FindByEmail implements [Repository].
func (r *repository) FindByEmail(ctx context.Context, email string) (User, error) {
	var err error
	var user User

	err = r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *repository) Update(ctx context.Context, user User) (User, error) {
	var err error

	err = r.DB.WithContext(ctx).Save(&user).Error
	return user, err
}

// Create implements [Repository].
func (r *repository) Create(ctx context.Context, user User) (User, error) {
	var err error

	err = r.DB.WithContext(ctx).Create(&user).Error

	return user, err

}

// Delete implements [Repository].
func (r *repository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&User{}, id).Error
}

// FindByID implements [Repository].
func (r *repository) FindByID(ctx context.Context, id uint) (User, error) {
	var err error
	var user User

	err = r.DB.WithContext(ctx).First(&user, id).Error

	return user, err

}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}
