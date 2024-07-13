package routers

import (
	"frz/controllers/admin"
	"frz/tools"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.POST("/validAuth", admin.ValidAuth)
		adminRouters.POST("/firstValid", tools.AuthMiddleware, admin.FirstValid)
		adminRouters.POST("/examineLeaveMessage", tools.AuthMiddleware, admin.ExamineLeaveMessage)
		adminRouters.POST("/uploadBlogInfo", tools.AuthMiddleware, admin.UploadBlogInfo)
		adminRouters.POST("/sendEmail", tools.AuthMiddleware, admin.UpdateInfo)
	}
}
