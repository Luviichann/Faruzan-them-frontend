package messageboard

import (
	"errors"
	"fmt"
	"frz/models"
	"frz/tools"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"gorm.io/gorm"
)

func SendEmail(ctx *gin.Context) {
	mailInfo := struct {
		Email string `json:"email"`
		Type  string `json:"type"`
		Key   string `json:"key"`
	}{}
	if err := ctx.ShouldBind(&mailInfo); err != nil {
		tools.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "错误")
		return
	}
	var ec models.EmailCode
	models.DB.Where("email = ? AND type = ?", mailInfo.Email, mailInfo.Type).Delete(&ec)
	e := email.NewEmail()
	//设置发送方的邮箱
	// e.From = "Faruzan <959994590@qq.com>"
	e.From = fmt.Sprintf("Faruzan <%s>", models.E_USERNAME)
	// 设置接收方的邮箱
	e.To = []string{mailInfo.Email}
	//设置主题
	e.Subject = "邮箱验证码"
	//设置文件发送的内容
	code := GenerateRandomString(10)
	e.HTML = []byte(fmt.Sprintf(`
	<p>你的验证码是%s，5分钟内有效。不要告诉别人哦。</p>
        <style>
            p{
                color: skyblue;
            }
        </style>
	`, code))
	//设置服务器相关的配置
	err := e.Send(models.E_ADDR, smtp.PlainAuth("", models.E_USERNAME, models.E_PASSWORD, models.E_HOST))
	if err != nil {
		log.Fatal(err)
		return
	}
	tools.Success(ctx, gin.H{
		"code": code,
	}, "邮箱验证码")
	emailCode := models.EmailCode{
		Email: mailInfo.Email,
		Code:  code,
		Type:  mailInfo.Type,
	}
	models.DB.Create(&emailCode)
	// 一段时间以后验证码自动失效。
	// time.Sleep(time.Second * 20)
	// models.DB.Where("email = ? AND code = ? AND type = ?", emailCode.Email, emailCode.Code, emailCode.Type).Delete(&emailCode)
}

// 生成验证码。
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 验证验证码正确性。
func ValidEmailCode(ctx *gin.Context) {
	emailCode := models.EmailCode{}
	if err := ctx.ShouldBind(&emailCode); err != nil {
		tools.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "错误")
		return
	}
	// 在数据库中查找是否有这么一组匹配的邮箱和验证码
	var e models.EmailCode
	if err := models.DB.Where("email = ? AND code = ? AND type = ?", emailCode.Email, emailCode.Code, emailCode.Type).First(&e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没找到。
			tools.Fail(ctx, gin.H{}, "验证码错误")
			return
		}
		tools.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "错误")
		return
	}
	tools.Success(ctx, gin.H{}, "验证码正确")
	models.DB.Where("email = ? AND code = ? AND type = ?", emailCode.Email, emailCode.Code, emailCode.Type).Delete(&e)
}
