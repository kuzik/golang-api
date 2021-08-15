package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/connection"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
	"gitlab.com/url-builder/go-admin/src/router"
)

func main() {
	// Load configs
	config.Setup()

	r := gin.New()
	router.Register(r)
	db, err := connection.Connect()
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(db)
		repositories.RegisterRepositories(db)
		fmt.Println("Database successfully connected")
	}

	r.Run()
}
