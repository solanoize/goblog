package repositories

import "gorm.io/gorm"

type PaginatorRepository interface {
	Apply(db *gorm.DB, page int, limit int) func(db *gorm.DB) *gorm.DB
}

type paginatorRepository struct {
	MaxLimit int
}

// Apply implements [PaginatorRepository].
func (p *paginatorRepository) Apply(db *gorm.DB, page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Validasi biar nggak error kalau user masukin angka ngaco
		if page <= 0 {
			page = 1
		}

		switch {
		case limit > p.MaxLimit:
			limit = p.MaxLimit // Batasin maksimal "max limit" data per request
		case limit <= 0:
			limit = 10 // Default 10 data kalau ga diisi
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func NewPaginatorRepository(maxLimit int) PaginatorRepository {
	return &paginatorRepository{MaxLimit: maxLimit}
}
