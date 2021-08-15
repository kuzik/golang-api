package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	Connection *gorm.DB
}

func (r Repository) paginate(page int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(database *gorm.DB) *gorm.DB {
		return database.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func (r Repository) connect(connection *gorm.DB) {
	r.Connection = connection
}

func RegisterRepositories(connection *gorm.DB) {
	AuthRepository = authRepository{Repository{Connection: connection}}
	UrlRepository = urlRepository{Repository{Connection: connection}}
}
