package apiv1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
	request "gitlab.com/url-builder/go-admin/src/requests/apiv1"
	"net/http"
	"strconv"
)

// ListUrl
// @Summary Get urls list
// @Tag url
// @Produce  json
// @Success 200 {object} []models.Url
// @Security ApiKeyAuth
// @Router /api/v1/url [get]
func ListUrl(context *gin.Context) {

	context.JSON(200, repositories.UserRepository.All())
}

// ViewUrl
// @Summary Get a single url
// @Tag url
// @Accept json
// @Produce  json
// @Param id path int true "Url ID"
// @Success 200 {object} request.UrlRequest
// @Security ApiKeyAuth
// @Router /api/v1/url/{id} [get]
func ViewUrl(context *gin.Context) {
	urlId, _ := strconv.Atoi(context.Param("id"))

	context.JSON(200, repositories.UserRepository.Find(urlId))
}

// CreateUrl
// @Summary Create new url
// @Accept  json
// @Produce  json
// @Param url body request.UrlRequest true "URL label"
// @Success 200 {object} request.UrlRequest
// @Security ApiKeyAuth
// @Router /api/v1/url/ [post]
func CreateUrl(context *gin.Context) {
	var urlInfo request.UrlRequest
	if err := context.ShouldBind(&urlInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlEntity := repositories.UserRepository.Create(urlInfo)

	context.JSON(http.StatusCreated, urlEntity)
}

// UpdateUrl
// @Summary Update an existing url
// @Accept  json
// @Produce  json
// @Param id path int true "Url ID"
// @Param url body request.UrlRequest true "URL label"
// @Success 200 {object} request.UrlRequest
// @Security ApiKeyAuth
// @Router /api/v1/url/{id} [put]
func UpdateUrl(context *gin.Context) {
	var urlInfo request.UrlRequest
	if err := context.ShouldBind(&urlInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlEntity := repositories.UserRepository.Update(prepareId(context, "id"), urlInfo)

	context.JSON(http.StatusOK, urlEntity)
}

// DeleteUrl
// @Summary Delete an existing url
// @Produce  json
// @Param id path int true "Url ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /api/v1/url/{id} [delete]
func DeleteUrl(context *gin.Context) {
	repositories.UserRepository.Delete(prepareId(context, "id"))

	context.JSON(http.StatusOK, "ok")
}

// Helper method for preparing ids extracted from url
func prepareId(context *gin.Context, name string) int {
	urlId, _ := strconv.Atoi(context.Param(name))

	return urlId
}
