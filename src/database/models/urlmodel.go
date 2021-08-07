package modelsimport

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Label       string
	Destination string
}
