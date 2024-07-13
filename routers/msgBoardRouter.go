package routers

import (
	"frz/controllers/messageboard"

	"github.com/gin-gonic/gin"
)

func MsgBoardRouter(r *gin.Engine) {
	msgBoardRouters := r.Group("/msgBoard")
	{
		// 发送邮件验证。
		msgBoardRouters.POST("/sendEmail", messageboard.SendEmail)
		// 验证邮箱验证码正确性。
		msgBoardRouters.POST("/validEmailCode", messageboard.ValidEmailCode)
		// 添加留言。
		msgBoardRouters.POST("/addLeaveMessage", messageboard.AddLeaveMessage)
		// 查询留言。
		msgBoardRouters.GET("/getLeaveMessage/:examine", messageboard.GetLeaveMessage)
	}
}
