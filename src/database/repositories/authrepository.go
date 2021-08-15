package repositories

import (
	"gitlab.com/url-builder/go-admin/src/database/models"
	"gorm.io/gorm"
)

type authRepository struct {
	Repository
}

var AuthRepository authRepository

// CheckAuth checks if authentication information exists
func (r authRepository) CheckAuth(username, password string) (int, error) {
	var auth models.User
	err := r.Repository.Connection.Where(models.User{Username: username, Password: password}).First(&auth).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if auth.ID > 0 {
		return auth.ID, nil
	}

	return 0, nil
}
