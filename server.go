package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
	"gitlab.com/url-builder/go-admin/src/router"
)

func main() {
	// Load configs
	config.Setup()

	r := gin.New()
	router.Register(r)
	_, err := repositories.Connect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database successfully connected")

	r.Run()
}
