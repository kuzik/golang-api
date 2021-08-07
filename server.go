package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/router"
)

func main() {
	r := gin.New()
	router.InitAPI(r)
	r.Run()
}
