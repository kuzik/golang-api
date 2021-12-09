package repositories

import (
	"gitlab.com/url-builder/go-admin/src/database/models"
	request "gitlab.com/url-builder/go-admin/src/requests/apiv1"
)

type urlRepository struct {
	Repository
}

type URLRepository interface {
	Create(url request.URLRequest) models.Url
	Update(urlID int, urlRequest request.URLRequest) models.Url
	All(page, pageSize int) []models.Url
	Find(urlID int) models.Url
	Delete(urlID int) bool
}

func (r urlRepository) Create(url request.URLRequest) models.Url {
	urlEntity := models.UrlFromRequest(url)

	r.Connection.Create(&urlEntity)

	return urlEntity
}

func (r urlRepository) Update(urlID int, urlRequest request.URLRequest) models.Url {
	var urlEntity models.Url

	r.Connection.First(&urlEntity, urlID)

	r.Connection.Model(&urlEntity).Updates(models.Url{Label: urlRequest.Label, Destination: urlRequest.Destination})

	return urlEntity
}

func (r urlRepository) All(page, pageSize int) []models.Url {
	var urls []models.Url

	r.Connection.Scopes(r.paginate(page, pageSize)).Find(&urls)

	return urls
}

func (r urlRepository) Find(urlID int) models.Url {
	var urlEntity models.Url

	r.Connection.First(&urlEntity, urlID)

	return urlEntity
}

func (r urlRepository) Delete(urlID int) bool {
	var urlEntity models.Url

	r.Connection.Delete(&urlEntity, urlID)

	return true
}
