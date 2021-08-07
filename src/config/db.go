package config

import "os"

type db struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Prefix   string
}

var Db = &db{}

func setupDb() {
	Db.Type = os.Getenv("DATABASE_TYPE")
	Db.User = os.Getenv("DATABASE_USER")
	Db.Password = os.Getenv("DATABASE_PASSWORD")
	Db.Host = os.Getenv("DATABASE_HOST")
	Db.Port = os.Getenv("DATABASE_PORT")
	Db.Name = os.Getenv("DATABASE_NAME")
	Db.Prefix = os.Getenv("DATABASE_PREFIX")
}
