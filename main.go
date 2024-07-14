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
		c.Writer.Header().Set("Access-Control-Allow-Origin", models.CORS)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
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
		ctx.String(http.StatusOK, "Hello World!")
	})
	routers.IndexRouter(r)
	routers.MsgBoardRouter(r)
	routers.SubscribeRouter(r)
	routers.AdminRouter(r)
	r.Run(":8820")
}
