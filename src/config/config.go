package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Setup() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	setupDb()
	setupAppConfig()
}
