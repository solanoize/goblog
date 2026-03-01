package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	UserID uint  `gorm:"index"`
	User   *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Title   string `gorm:"type:varchar(200);not null"`
	Slug    string `gorm:"type:varchar(200);uniqueIndex;not null"`
	Content string `gorm:"type:text;not null"`

	Published bool `gorm:"default:false"`
}
