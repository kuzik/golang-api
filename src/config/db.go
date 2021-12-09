package config

import (
	"fmt"
	"os"
)

type db struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Prefix   string
}

var DB = &db{}

func setupDB() {
	DB.Type = os.Getenv("DATABASE_TYPE")
	DB.User = os.Getenv("DATABASE_USER")
	DB.Password = os.Getenv("DATABASE_PASSWORD")
	DB.Host = os.Getenv("DATABASE_HOST")
	DB.Port = os.Getenv("DATABASE_PORT")
	DB.Name = os.Getenv("DATABASE_NAME")
	DB.Prefix = os.Getenv("DATABASE_PREFIX")
}

// AsDsn
// Helper function for serializing database configuration to DSN string
func (config db) AsDsn() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}
