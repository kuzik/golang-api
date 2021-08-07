package apiv1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
)

// ListUrl
// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Router /apiv1/v1/articles/{id} [get]
func ListUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4", config.Db.Host})
}

func ViewUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}

func CreateUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}

func UpdateUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}

func DeleteUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
}
