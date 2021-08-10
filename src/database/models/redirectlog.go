package models

type RedirectLog struct {
	Model
	UrlID       int
	ClientIp    string `gorm:"<-:create;size:16"`
	Referer     string
	CountryCode string

	Url Url
}
