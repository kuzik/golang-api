package router

import (
	"github.com/gin-gonic/gin"
	apiController "gitlab.com/url-builder/go-admin/src/controllers/api"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "gitlab.com/url-builder/go-admin/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is documentation for redirect engine admin API

// InitAPI @host 127.0.0.1:8080
// @BasePath /v1
func InitAPI(router *gin.Engine) {
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := router.Group("/api/v1")
	{
		api.GET("/url/", apiController.ListUrl)
		api.GET("/url/:id", apiController.ViewUrl)
		api.POST("/url/:id", apiController.UpdateUrl)
		api.POST("/url/", apiController.CreateUrl)
		api.DELETE("/url/:id", apiController.DeleteUrl)
	}
}
