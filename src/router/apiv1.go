package router

import (
	"github.com/gin-gonic/gin"

	"gitlab.com/url-builder/go-admin/src/controllers/apiv1"
	"gitlab.com/url-builder/go-admin/src/middleware"
)

func registerAPI(router *gin.Engine) {
	api := router.Group("/api/v1")
	// Add jwt Authorization
	api.Use(middleware.JWT())
	{
		api.GET("/url/", apiv1.ListUrl)
		api.GET("/url/:id", apiv1.ViewUrl)
		api.PUT("/url/:id", apiv1.UpdateUrl)
		api.POST("/url/", apiv1.CreateUrl)
		api.DELETE("/url/:id", apiv1.DeleteUrl)
	}
}
