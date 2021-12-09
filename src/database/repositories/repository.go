package repositories

import (
	"gorm.io/gorm"
)

type RepositoryPool interface {
	GetAuthRepository() AuthRepository
	GetURLRepository() URLRepository
}

type Repository struct {
	Connection *gorm.DB
}

type repositoryPool struct {
	authRepository AuthRepository
	urlRepository  URLRepository
}

func (p repositoryPool) GetAuthRepository() AuthRepository {
	return p.authRepository
}

func (p repositoryPool) GetURLRepository() URLRepository {
	return p.urlRepository
}

func (r Repository) paginate(page int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(database *gorm.DB) *gorm.DB {
		return database.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func GetRepositoriesPool(connection *gorm.DB) RepositoryPool {
	return repositoryPool{
		authRepository{Repository{Connection: connection}},
		urlRepository{Repository{Connection: connection}},
	}
}
