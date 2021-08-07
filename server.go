package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/router"
)

func main() {
	// Load configs
	config.Setup()

	r := gin.New()
	router.InitAPI(r)
	r.Run()
}
