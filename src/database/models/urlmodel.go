package models

import (
	"gitlab.com/url-builder/go-admin/src/requests/apiv1"
)

type Url struct {
	Model
	UserID      int
	CampaignID  int
	Label       string
	Destination string

	User     User
	Campaign Campaign
}

func UrlFromRequest(request apiv1.UrlRequest) Url {
	url := Url{
		Label:       request.Label,
		Destination: request.Destination,
	}

	return url
}
