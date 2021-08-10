package repositories

import "gorm.io/gorm"

type Repository struct {
}

func (r Repository) paginate(page int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(database *gorm.DB) *gorm.DB {
		return database.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
