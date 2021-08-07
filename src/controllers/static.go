package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WelcomePage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
