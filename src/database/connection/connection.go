package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/models"
)

func Connect(dbConfig config.DB) (*gorm.DB, error) {
	db, mysqlErr := gorm.Open(mysql.Open(dbConfig.AsDsn()), &gorm.Config{})

	return db, mysqlErr
}

// Migrate
// For performance improvements we run migration from dedicated command only
func Migrate(db *gorm.DB) error {
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
