package controllers

import (
	"time"

	"gitlab.com/url-builder/go-admin/src/database/repositories"
	"gitlab.com/url-builder/go-admin/src/services"

	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.com/url-builder/go-admin/src/requests"
)

type authController struct {
	authRepository  repositories.AuthRepository
	securityService services.SecurityService
}

func registerAuth(router *gin.Engine, authRepository repositories.AuthRepository, securityService services.SecurityService) {
	controller := authController{
		authRepository:  authRepository,
		securityService: securityService,
	}

	router.POST("/auth", controller.GetAuth)
}

// GetAuth @Summary Get Auth
// @Produce  json
// @Param url body requests.AuthRequest true "Auth data"
// @Success 200 {object} interface{}
// @Router /auth [post]
func (a authController) GetAuth(context *gin.Context) {
	var authForm requests.AuthRequest
	if err := context.ShouldBind(&authForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := a.securityService.EncodeSha(authForm.Password)
	userID, err := a.authRepository.CheckAuth(authForm.Username, password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if userID == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	token, err := a.securityService.GenerateToken(authForm.Username, userID, expireTime)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
