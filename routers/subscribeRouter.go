package routers

import (
	"frz/controllers/subscribeto"

	"github.com/gin-gonic/gin"
)

func SubscribeRouter(r *gin.Engine) {
	subscribeRouters := r.Group("/subscribe")
	{
		subscribeRouters.POST("/option", subscribeto.Option)
	}
}
