package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
	"gitlab.com/url-builder/go-admin/src/services"
)

func Register(router *gin.Engine, pool repositories.RepositoryPool, config config.Config) {
	securityService := services.SecurityService{Secret: config.App.Secret}

	registerStaticRoutes(router)
	registerAPI(router, pool.GetURLRepository(), securityService)
	registerSwagger(router)

	registerAuth(
		router,
		pool.GetAuthRepository(),
		securityService,
	)
}
