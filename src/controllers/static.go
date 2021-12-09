package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomePage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
