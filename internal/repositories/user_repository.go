package repositories

import (
	"github.com/solanoize/goblog/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(page int, limit int, search string) ([]models.User, int64, error)
	FindById(id uint) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	DB                  *gorm.DB
	PaginatorRepository PaginatorRepository
}

// Create implements [UserRepository].
func (u *userRepository) Create(user models.User) (models.User, error) {
	var err error
	err = u.DB.Create(&user).Error
	return user, err
}

// Delete implements [UserRepository].
func (u *userRepository) Delete(id uint) error {
	return u.DB.Delete(&models.User{}, id).Error
}

// FindAll implements [UserRepository].
func (u *userRepository) FindAll(page int, limit int, search string) ([]models.User, int64, error) {
	var err error
	var query *gorm.DB
	var totalCount int64
	var searchParams string
	var users []models.User

	query = u.DB.Model(&models.User{})

	if search != "" {
		searchParams = "%" + search + "%"
		query = query.Where("username LIKE ? OR email LIKE ?", searchParams, searchParams)
	}

	query.Count(&totalCount)

	err = query.Scopes(u.PaginatorRepository.Apply(u.DB, page, limit)).Find(&users).Error

	return users, totalCount, err
}

// FindByEmail implements [UserRepository].
func (u *userRepository) FindByEmail(email string) (models.User, error) {
	var err error
	var user models.User

	err = u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// FindById implements [UserRepository].
func (u *userRepository) FindById(id uint) (models.User, error) {
	var err error
	var user models.User

	err = u.DB.First(&user, id).Error
	return user, err
}

// Update implements [UserRepository].
func (u *userRepository) Update(user models.User) (models.User, error) {
	err := u.DB.Save(&user).Error
	return user, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB:                  db,
		PaginatorRepository: NewPaginatorRepository(100),
	}
}
