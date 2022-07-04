package routers

import (
	"gin-api/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		// 2.多文件夹r.LoadHTMLGlob("template/**/*")模板配置
		apiRouter.GET("/", api.ApiController{}.Index)
		apiRouter.GET("/list", api.ApiController{}.List)
		apiRouter.GET("/post", api.ApiController{}.Post)
	}
}
