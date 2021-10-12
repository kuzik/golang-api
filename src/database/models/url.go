package models

import (
	"gitlab.com/url-builder/go-admin/src/requests/apiv1"
)

type Url struct {
	Model
	UserID      int    `gorm:"column:user_id"`
	Label       string `gorm:"column:label"`
	Destination string `gorm:"column:destination"`

	User      User
	Campaigns []Campaign `gorm:"many2many:url_campaigns;"`
}

func UrlFromRequest(request apiv1.UrlRequest) Url {
	url := Url{
		Label:       request.Label,
		Destination: request.Destination,
		UserID:      request.UserId,
	}

	return url
}
