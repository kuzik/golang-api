package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App App
	DB  DB
}

type App struct {
	Secret  string
	RootDir string
}

type DB struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	NameTest string
	Prefix   string
}

// AsDsn
// Helper function for serializing database configuration to DSN string
func (config DB) AsDsn() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}

// AsDsnNoDb
// Helper function for serializing database configuration to DSN string
func (config DB) AsDsnNoDb() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
}

func LoadConfigs(path string) Config {
	err := godotenv.Load(path)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		DB: DB{
			Type:     os.Getenv("DATABASE_TYPE"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Name:     os.Getenv("DATABASE_NAME"),
			Prefix:   os.Getenv("DATABASE_PREFIX"),
		},
		App: App{
			Secret:  os.Getenv("APP_SECRET"),
			RootDir: os.Getenv("ROOT_DIR"),
		},
	}
}
