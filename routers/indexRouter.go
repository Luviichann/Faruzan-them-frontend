package routers

import (
	"frz/controllers/index"

	"github.com/gin-gonic/gin"
)

func IndexRouter(r *gin.Engine) {
	indexRouters := r.Group("/index")
	{
		indexRouters.GET("/blogs/:examine", index.GetBlogs)
	}
}
