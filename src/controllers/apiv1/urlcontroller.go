package apiv1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/database/models"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
	request "gitlab.com/url-builder/go-admin/src/requests/apiv1"
	"net/http"
)

// ListUrl
// @Summary Get a single article
// @Tag url
// @Produce  json
// @Success 200 {object} []models.Url
// @Router /api/v1/url [get]
func ListUrl(context *gin.Context) {

	context.JSON(200, repositories.UserRepository.All())
}

// ViewUrl
// @Summary Get a single article
// @Tag url
// @Accept json
// @Produce  json
// @Param  id path int true "Bottle ID"
// @Router /api/v1/url/{id} [get]
func ViewUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}

// CreateUrl
// @Summary Create new url
// @Accept  json
// @Produce  json
// @Param url body request.CreateUrlRequest true "URL label"
// @Success 200 {object} request.CreateUrlRequest
// @Router /api/v1/url/ [post]
func CreateUrl(context *gin.Context) {
	var urlInfo request.CreateUrlRequest
	if err := context.ShouldBind(&urlInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlEntity := models.UrlFromRequest(urlInfo)
	repositories.UserRepository.Create(&urlEntity)

	context.JSON(http.StatusCreated, urlEntity)
}

func UpdateUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}

func DeleteUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}
