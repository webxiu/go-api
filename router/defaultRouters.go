package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouters(r *gin.Engine) {
	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", func(c *gin.Context) {
			page := c.Query("page")
			pageSize := c.DefaultQuery("pageSize", "10")
			c.JSON(http.StatusOK, gin.H{
				"name":     "张三",
				"page":     page,
				"pageSize": pageSize,
			})
		})

		// 2.多文件夹r.LoadHTMLGlob("template/**/*")模板配置
		defaultRouter.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home/index.html", gin.H{
				"title": "首页数据",
				"time":  1656233903,
			})
		})
		defaultRouter.GET("/goods", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home/goods.html", gin.H{
				"title": "首页数据goods",
			})
		})
	}
}
