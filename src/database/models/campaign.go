package models

import "time"

type Campaign struct {
	Model
	Name      string
	UserID    int
	StartDate *time.Time
	EndDate   *time.Time

	User User
}
