package repositories

import (
	"gitlab.com/url-builder/go-admin/src/database/models"
	request "gitlab.com/url-builder/go-admin/src/requests/apiv1"
)

type urlRepository struct{}

var UserRepository = urlRepository{}

func (r urlRepository) Create(url request.UrlRequest) models.Url {
	urlEntity := models.UrlFromRequest(url)

	db.Create(&urlEntity)

	return urlEntity
}

func (r urlRepository) Update(urlId int, urlRequest request.UrlRequest) models.Url {
	var urlEntity models.Url

	db.First(&urlEntity, urlId)

	db.Model(&urlEntity).Updates(models.Url{Label: urlRequest.Label, Destination: urlRequest.Destination})

	return urlEntity
}

func (r urlRepository) All() []models.Url {
	var urls []models.Url

	db.Limit(20).Find(&urls)

	return urls
}

func (r urlRepository) Find(urlId int) models.Url {
	var urlEntity models.Url

	db.First(&urlEntity, urlId)

	return urlEntity
}

func (r urlRepository) Delete(urlId int) bool {
	var urlEntity models.Url

	db.Delete(&urlEntity, urlId)

	return true
}
