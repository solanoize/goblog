package repositories

import (
	"github.com/solanoize/goblog/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(page int, limit int, search string) ([]models.User, int64, error)
	Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB
	FindById(id uint) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func (u *userRepository) Paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Validasi biar nggak error kalau user masukin angka ngaco
		if page <= 0 {
			page = 1
		}

		switch {
		case limit > 100:
			limit = 100 // Batasin maksimal 100 data per request biar server ga jebol
		case limit <= 0:
			limit = 10 // Default 10 data kalau ga diisi
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
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
	var query *gorm.DB = u.DB.Model(&models.User{})
	var paginate func(db *gorm.DB) *gorm.DB = func(db *gorm.DB) *gorm.DB {
		// Validasi biar nggak error kalau user masukin angka ngaco
		if page <= 0 {
			page = 1
		}

		switch {
		case limit > 100:
			limit = 100 // Batasin maksimal 100 data per request biar server ga jebol
		case limit <= 0:
			limit = 10 // Default 10 data kalau ga diisi
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}

	if search != "" {
		var searchParams string = "%" + search + "%"
		query = query.Where("username LIKE ? OR email LIKE ?", searchParams, searchParams)
	}

	var totalCount int64
	query.Count(&totalCount)

	var err error
	var users []models.User

	err = query.Scopes(paginate).Find(&users).Error

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
	return &userRepository{DB: db}
}
