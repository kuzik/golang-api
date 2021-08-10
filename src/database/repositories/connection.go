package repositories

import (
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (*gorm.DB, error) {
	var mysqlErr error
	db, mysqlErr = gorm.Open(mysql.Open(config.Db.AsDsn()), &gorm.Config{})

	err := db.AutoMigrate(
		&models.Url{},
		&models.User{},
		&models.RedirectLog{},
		&models.Campaign{},
	)
	if err != nil {
		return nil, err
	}

	return db, mysqlErr
}
