package models

import (
	"fmt"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	UserID uint  `gorm:"index"`
	User   *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Title   string `gorm:"type:varchar(200);not null"`
	Slug    string `gorm:"type:varchar(200);uniqueIndex;not null"`
	Content string `gorm:"type:text;not null"`

	Published bool `gorm:"default:false"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 1. Generate slug dasar
	baseSlug := slug.Make(p.Title)
	finalSlug := baseSlug
	counter := 1

	// 2. Loop buat ngecek apakah slug udah dipake
	for {
		var count int64
		// Kita pake 'tx' (transaction) bawaan hook biar konsisten
		// Cek apakah ada post lain yang punya slug yang sama
		err := tx.Model(&Post{}).Where("slug = ?", finalSlug).Count(&count).Error
		if err != nil {
			return err
		}

		if count == 0 {
			// Slug aman! Belum ada yang pake
			break
		}

		// Kalau ada yang pake, tambahin angka di belakang (misal: judul-post-1)
		finalSlug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}

	p.Slug = finalSlug
	return nil
}
