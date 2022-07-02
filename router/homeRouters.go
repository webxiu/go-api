package routers

import (
	"gin-api/controllers/home"

	"github.com/gin-gonic/gin"
)

func HomeRouters(r *gin.Engine) {
	defaultRouter := r.Group("/")
	{
		// 2.多文件夹r.LoadHTMLGlob("template/**/*")模板配置
		defaultRouter.GET("/index", home.HomeController{}.Index)
		defaultRouter.GET("/page", home.HomeController{}.Page)
		defaultRouter.GET("/goods", home.HomeController{}.Goods)

	}
}
