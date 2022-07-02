package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		// 2.多文件夹r.LoadHTMLGlob("template/**/*")模板配置
		apiRouter.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home/index.html", gin.H{
				"title": "首页数据",
				"time":  1656233903,
			})
		})
		apiRouter.GET("/api/list", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home/goods.html", gin.H{
				"title": "api/list",
			})
		})
		// 1.获取参数
		apiRouter.POST("/post", func(c *gin.Context) {
			username := c.PostForm("username") // 获取表单提交值
			pageSize := c.DefaultPostForm("pageSize", "10")
			c.JSON(http.StatusOK, gin.H{
				"name":     "张三",
				"username": username,
				"pageSize": pageSize,
			})
		})
	}
}
