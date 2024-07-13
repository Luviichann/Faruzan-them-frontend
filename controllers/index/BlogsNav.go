package index

import (
	"frz/models"
	"frz/tools"

	"github.com/gin-gonic/gin"
)

func GetBlogs(ctx *gin.Context) {
	examine := ctx.Param("examine")
	blogsList := []models.Blog{}
	models.DB.Where("examine = ?", examine).Find(&blogsList)
	tools.Success(ctx, gin.H{
		"blogsList": blogsList,
	}, "博客列表")
}
