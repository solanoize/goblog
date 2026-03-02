package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	IsAdmin  bool   `gorm:"default:false"`
	IsStaff  bool   `gorm:"default:false"`
	Active   bool   `gorm:"default:true"`
}
