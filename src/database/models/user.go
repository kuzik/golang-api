package models

type User struct {
	ID       int    `gorm:"primary_key;column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password;->:false" json:"password"`
}
