package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type staticController struct {
}

func registerStaticRoutes(router *gin.Engine) {
	controller := staticController{}

	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.WelcomePage)
}

func (s staticController) WelcomePage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
