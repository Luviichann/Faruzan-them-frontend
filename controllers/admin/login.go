package admin

import (
	"fmt"
	"frz/models"
	"frz/tools"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidAuth(ctx *gin.Context) {
	ap := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}
	if err := ctx.ShouldBind(&ap); err != nil {
		tools.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "错误")
		return
	}
	if ap.Account != models.ACCOUNT || ap.Password != models.PASSWORD {
		return
	}
	token, err := tools.ReleaseToken(ap.Account)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token get error:%v", err)
		return
	}
	tools.Success(ctx, gin.H{
		"token": token,
	}, "super")
}
