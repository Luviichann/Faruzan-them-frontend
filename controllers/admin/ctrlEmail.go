package admin

import (
	"fmt"
	"frz/models"
	"frz/tools"
	"log"
	"net/smtp"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

// 更新消息发布
func UpdateInfo(ctx *gin.Context) {
	ei := struct {
		Content string `json:"content"`
		Key     string `json:"key"`
	}{}
	if err := ctx.ShouldBind(&ei); err != nil {
		tools.Fail(ctx, gin.H{}, "未知错误")
		return
	}
	mailList := []models.SubMail{}
	models.DB.Find(&mailList)
	to := []string{}
	for i := 0; i < len(mailList); i++ {
		to = append(to, mailList[i].Email)
	}
	sendEmail(ei.Content, to)
}

func sendEmail(content string, receive []string) {
	fmt.Printf("content: %v\n", content)
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = fmt.Sprintf("Faruzan <%s>", models.E_USERNAME)
	// 设置接收方的邮箱
	e.To = receive
	//设置主题
	e.Subject = "更新消息"
	//设置文件发送的内容
	e.HTML = []byte(content)
	//设置服务器相关的配置
	err := e.Send(models.E_ADDR, smtp.PlainAuth("", models.E_USERNAME, models.E_PASSWORD, models.E_HOST))
	if err != nil {
		log.Fatal(err)
		return
	}
}
