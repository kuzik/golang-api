package router

import (
	"github.com/gin-gonic/gin"

	"gitlab.com/url-builder/go-admin/src/controllers"
)

func registerStaticRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.WelcomePage)
}
