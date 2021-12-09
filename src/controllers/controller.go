package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
)

func Register(router *gin.Engine, pool repositories.RepositoryPool) {
	registerStaticRoutes(router)
	registerAPI(router, pool.GetURLRepository())
	registerSwagger(router)
	registerAuth(router, pool.GetAuthRepository())
}
