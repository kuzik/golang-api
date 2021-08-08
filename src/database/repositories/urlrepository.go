package repositories

import (
	"gitlab.com/url-builder/go-admin/src/database/models"
	"gorm.io/gorm"
)

type urlRepository struct{}

var UserRepository = urlRepository{}

func (r urlRepository) Create(url *models.Url) *gorm.DB {
	return db.Create(url)
}

func (r urlRepository) All() interface{} {
	var urls []models.Url

	db.Limit(20).Find(&urls)

	return urls
}
