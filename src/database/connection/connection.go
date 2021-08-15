package connection

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

	return db, mysqlErr
}

// Migrate
// For performance improvements we run migration from dedicated command only
func Migrate() error {

	err := db.AutoMigrate(
		&models.Url{},
		&models.User{},
		&models.RedirectLog{},
		&models.Campaign{},
	)
	if err != nil {
		return err
	}

	return nil
}
