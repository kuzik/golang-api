package repositories

import (
	"gitlab.com/url-builder/go-admin/src/database/models"
	request "gitlab.com/url-builder/go-admin/src/requests/apiv1"
)

type urlRepository struct {
	Repository
}

var UrlRepository urlRepository

func (r urlRepository) Create(url request.UrlRequest) models.Url {
	urlEntity := models.UrlFromRequest(url)

	r.Connection.Create(&urlEntity)

	return urlEntity
}

func (r urlRepository) Update(urlId int, urlRequest request.UrlRequest) models.Url {
	var urlEntity models.Url

	r.Connection.First(&urlEntity, urlId)

	r.Connection.Model(&urlEntity).Updates(models.Url{Label: urlRequest.Label, Destination: urlRequest.Destination})

	return urlEntity
}

func (r urlRepository) All(page, pageSize int) []models.Url {
	var urls []models.Url

	r.Connection.Scopes(r.paginate(page, pageSize)).Find(&urls)

	return urls
}

func (r urlRepository) Find(urlId int) models.Url {
	var urlEntity models.Url

	r.Connection.First(&urlEntity, urlId)

	return urlEntity
}

func (r urlRepository) Delete(urlId int) bool {
	var urlEntity models.Url

	r.Connection.Delete(&urlEntity, urlId)

	return true
}
