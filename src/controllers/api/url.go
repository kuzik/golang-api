package api

import (
	"github.com/gin-gonic/gin"
)

// UrlList
// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Router /api/v1/articles/{id} [get]
func ListUrl(context *gin.Context) {

	context.JSON(200, []string{"link1", "link2", "link3", "link4"})
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
