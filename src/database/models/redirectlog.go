package models

type RedirectLog struct {
	Model
	UrlID       int    `gorm:"column:url_id"`
	ClientIp    string `gorm:"column:client_ip;<-:create;size:16"`
	Referer     string `gorm:"column:referer"`
	CountryCode string `gorm:"column:country_code"`

	Url Url
}
