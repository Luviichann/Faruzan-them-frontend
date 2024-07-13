package messageboard

import (
	"encoding/json"
	"fmt"
	"frz/models"
	"frz/tools"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 添加留言
func AddLeaveMessage(ctx *gin.Context) {
	var data map[string]models.LeaveMessage
	if err := ctx.ShouldBind(&data); err != nil {
		tools.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "错误")
		return
	}
	// fmt.Printf("data: %v\n", data)
	leaveMsg := data["leaveMsg"]
	leaveMsg.Address = getAddress(leaveMsg.Ip)
	leaveMsg.Examine = "no"
	models.DB.Create(&leaveMsg)
}

// 获取留言
func GetLeaveMessage(ctx *gin.Context) {
	examine := ctx.Param("examine")
	leaveMsgList := []models.LeaveMessage{}
	if examine == "all" {
		models.DB.Find(&leaveMsgList)
	} else {
		models.DB.Where("examine = ?", examine).Find(&leaveMsgList)
	}
	tools.Success(ctx, gin.H{
		"leaveMsgList": leaveMsgList,
	}, "博客列表")
}

// 获取完整地址
func getAddress(ip string) string {
	url := fmt.Sprintf("https://opendata.baidu.com/api.php?query=%s&co=&resource_id=6006&oe=utf8", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "未知地址" + ip
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "未知地址" + ip
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "未知地址" + ip
	}
	var data map[string]any
	// 解析JSON响应
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "未知地址" + ip
	}

	location := data["data"].([]any)[0].(map[string]any)["location"]
	return getFirstPartBeforeSpace(location.(string))
}

// 获得地址的一部分
func getFirstPartBeforeSpace(input string) string {
	spaceIndex := strings.Index(input, " ")
	if spaceIndex == -1 {
		return input // 如果没有空格，返回整个字符串
	}
	return input[:spaceIndex]
}
