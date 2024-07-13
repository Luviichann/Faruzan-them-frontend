package admin

import (
	"frz/tools"

	"github.com/gin-gonic/gin"
)

func FirstValid(ctx *gin.Context) {
	tools.Success(ctx, gin.H{}, "success")
}
