package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gitlab.com/url-builder/go-admin/docs"
)

// @title URL management API
// @version 1.0
// @description This is documentation for redirect engine admin API
// @host 127.0.0.1:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func registerSwagger(router *gin.Engine) {
	swaggerGroup := router.Group("/", gin.BasicAuth(gin.Accounts{
		"api-admin": "1111",
	}))
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json") // The url pointing to API definition
	swaggerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
