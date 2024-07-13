package subscribeto

import (
	"errors"
	"frz/models"
	"frz/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 邮箱订阅。
func Option(ctx *gin.Context) {
	subInfo := struct {
		Email  string `json:"email"`
		Option string `json:"option"`
		Key    string `json:"key"`
	}{}
	if err := ctx.ShouldBind(&subInfo); err != nil {
		tools.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "错误")
		return
	}
	sm := models.SubMail{
		Email: subInfo.Email,
	}
	if subInfo.Option == "dyyt" {
		subscribe(sm, ctx)
	} else if subInfo.Option == "tvdy" {
		unsubscribe(sm, ctx)
	}
}

func subscribe(sm models.SubMail, ctx *gin.Context) {
	subMail := models.SubMail{}
	if err := models.DB.Where("email = ?", sm.Email).First(&subMail).Error; err == nil {
		tools.Fail(ctx, gin.H{}, "你已经定阅过了！")
		return
	}
	tools.Success(ctx, gin.H{}, "订阅成功！")
	models.DB.Create(&sm)
}

func unsubscribe(sm models.SubMail, ctx *gin.Context) {
	subMail := models.SubMail{}
	if err := models.DB.Where("email = ?", sm.Email).First(&subMail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没找到。
			tools.Fail(ctx, gin.H{}, "你还没有订阅！")
			return
		}
		return
	}
	tools.Success(ctx, gin.H{}, "退订成功！")
	models.DB.Where("email = ?", sm.Email).Delete(&sm)
}
