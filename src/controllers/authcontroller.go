package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/requests"
	authservice "gitlab.com/url-builder/go-admin/src/services/auth"
	"net/http"
)

// GetAuth @Summary Get Auth
// @Produce  json
// @Param url body requests.AuthRequest true "Auth data"
// @Success 200 {object} interface{}
// @Router /auth [post]
func GetAuth(context *gin.Context) {
	var authForm requests.AuthRequest
	if err := context.ShouldBind(&authForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := authservice.Auth{Username: authForm.Username, Password: authForm.Password}
	isExist, err := authService.Check()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if !isExist {
		context.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	token, err := authservice.GenerateToken(authForm.Username, authForm.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}