package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/controllers"
)

func registerAuth(router *gin.Engine) {
	router.POST("/auth", controllers.GetAuth)
}
