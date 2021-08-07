package router

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	registerAPI(router)
	registerStaticRoutes(router)
}
