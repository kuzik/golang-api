package models

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `gorm:"->:false" json:"password"`
}
