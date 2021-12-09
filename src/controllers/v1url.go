package controllers

import (
	"gitlab.com/url-builder/go-admin/src/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gitlab.com/url-builder/go-admin/src/database/repositories"
	request "gitlab.com/url-builder/go-admin/src/requests/apiv1"
)

type urlController struct {
	urlRepository repositories.URLRepository
}

func registerAPI(router *gin.Engine, urlRepository repositories.URLRepository) {
	controller := urlController{urlRepository: urlRepository}

	api := router.Group("/api/v1")
	// Add jwt Authorization
	api.Use(middleware.JWT())
	{
		api.GET("/url/", controller.ListURL)
		api.GET("/url/:id", controller.ViewURL)
		api.PUT("/url/:id", controller.UpdateURL)
		api.POST("/url/", controller.CreateURL)
		api.DELETE("/url/:id", controller.DeleteURL)
	}
}

// ListURL
// @Summary Get urls list
// @Tag url
// @Produce  json
// @Param page query int true "Current page" default(1)
// @Param page_size query int true "Page size" default(10)
// @Success 200 {object} []models.Url
// @Security ApiKeyAuth
// @Router /api/v1/url [get]
func (u urlController) ListURL(context *gin.Context) {

	page, _ := strconv.Atoi(context.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(context.Query("page_size"))
	if pageSize == 0 {
		pageSize = 10
	}

	context.JSON(200, u.urlRepository.All(page, pageSize))
}

// ViewURL
// @Summary Get a single url
// @Tag url
// @Accept json
// @Produce  json
// @Param id path int true "Url ID"
// @Success 200 {object} request.URLRequest
// @Security ApiKeyAuth
// @Router /api/v1/url/{id} [get]
func (u urlController) ViewURL(context *gin.Context) {
	urlID, _ := strconv.Atoi(context.Param("id"))

	context.JSON(200, u.urlRepository.Find(urlID))
}

// CreateURL
// @Summary Create new url
// @Accept  json
// @Produce  json
// @Param url body request.URLRequest true "URL label"
// @Success 200 {object} request.URLRequest
// @Security ApiKeyAuth
// @Router /api/v1/url/ [post]
func (u urlController) CreateURL(context *gin.Context) {
	var urlInfo request.URLRequest
	if err := context.ShouldBind(&urlInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urlInfo.UserID = context.MustGet("user_id").(int)

	urlEntity := u.urlRepository.Create(urlInfo)

	context.JSON(http.StatusCreated, urlEntity)
}

// UpdateURL
// @Summary Update an existing url
// @Accept  json
// @Produce  json
// @Param id path int true "Url ID"
// @Param url body request.URLRequest true "URL label"
// @Success 200 {object} request.URLRequest
// @Security ApiKeyAuth
// @Router /api/v1/url/{id} [put]
func (u urlController) UpdateURL(context *gin.Context) {
	var urlInfo request.URLRequest
	if err := context.ShouldBind(&urlInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urlInfo.UserID = context.MustGet("user_id").(int)

	urlEntity := u.urlRepository.Update(prepareId(context, "id"), urlInfo)

	context.JSON(http.StatusOK, urlEntity)
}

// DeleteURL
// @Summary Delete an existing url
// @Produce  json
// @Param id path int true "Url ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /api/v1/url/{id} [delete]
func (u urlController) DeleteURL(context *gin.Context) {
	u.urlRepository.Delete(prepareId(context, "id"))

	context.JSON(http.StatusOK, "ok")
}

// Helper method for preparing ids extracted from url
func prepareId(context *gin.Context, name string) int {
	urlID, _ := strconv.Atoi(context.Param(name))

	return urlID
}
