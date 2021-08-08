package models

import (
	"gitlab.com/url-builder/go-admin/src/requests/apiv1"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Label       string
	Destination string
}

func UrlFromRequest(request apiv1.CreateUrlRequest) Url {
	url := Url{
		Label:       request.Label,
		Destination: request.Destination,
	}

	return url
}
