package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
)

type staticController struct {
}

func registerStaticRoutes(router *gin.Engine, configs config.Config) {
	controller := staticController{}

	router.LoadHTMLGlob(configs.App.RootDir + "templates/*")
	router.GET("/", controller.WelcomePage)
}

func (s staticController) WelcomePage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
