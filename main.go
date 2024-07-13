package main

import (
	"frz/models"
	"frz/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cors := func(c *gin.Context) {
		// 允许特定的域进行跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", models.CORS)
		// 允许特定的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许特定的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带身份凭证（如Cookie）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		} else {
			// 继续处理请求
			c.Next()
		}

	}
	r.Use(cors)

	r.Static("/static/image", "./static/image")
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!") //http.StatusOK就等于200。
	})
	routers.IndexRouter(r)
	routers.MsgBoardRouter(r)
	routers.SubscribeRouter(r)
	routers.AdminRouter(r)
	r.Run(":8820") //默认端口是8000。括号里可以指定端口号，接收字符串类型，比如":4000"。注意":"不能省略！
}
