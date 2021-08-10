package repositories

import (
	"gitlab.com/url-builder/go-admin/src/database/models"
	"gorm.io/gorm"
)

type authRepository struct{}

var AuthRepository = authRepository{}

// CheckAuth checks if authentication information exists
func (r authRepository) CheckAuth(username, password string) (bool, error) {
	var auth models.User
	err := db.Select("id").Where(models.User{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
