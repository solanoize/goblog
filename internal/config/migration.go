package config

import (
	"github.com/solanoize/goblog/internal/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
