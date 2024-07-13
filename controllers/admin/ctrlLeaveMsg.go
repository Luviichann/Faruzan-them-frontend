package admin

import (
	"frz/models"
	"frz/tools"

	"github.com/gin-gonic/gin"
)

// 审核留言
func ExamineLeaveMessage(ctx *gin.Context) {
	// 接收id，examine是否通过。
	exInfo := struct {
		Id      int    `json:"id"`
		Examine string `json:"examine"`
	}{}
	if err := ctx.ShouldBind(&exInfo); err != nil {
		tools.Fail(ctx, gin.H{}, "未知错误")
		return
	}
	if exInfo.Examine == "yes" {
		// 审核通过，将展示。
		tools.Success(ctx, gin.H{}, "审核通过")
		var lm models.LeaveMessage
		lm.PassExamine(exInfo.Id)
	} else if exInfo.Examine == "no" {
		// 不通过删除。
		tools.Success(ctx, gin.H{}, "删除成功")
		var lm models.LeaveMessage
		lm.DeleteData(exInfo.Id)
	}
}
