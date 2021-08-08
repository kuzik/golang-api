package router

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	registerStaticRoutes(router)
	registerAPI(router)
	registerSwagger(router)
	registerAuth(router)
}
