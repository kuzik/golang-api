package models

import "time"

type Campaign struct {
	Model
	Name      string     `gorm:"column:name"`
	UserID    int        `gorm:"column:user_id"`
	StartDate *time.Time `gorm:"column:start_date"`
	EndDate   *time.Time `gorm:"column:end_date"`

	User User `db:"user"`
}
