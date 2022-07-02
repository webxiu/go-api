package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "首页数据",
		"time":  1656233903,
	})
}
func (con ApiController) List(c *gin.Context) {
	c.HTML(http.StatusOK, "home/goods.html", gin.H{
		"title": "api/list",
	})
}
func (con ApiController) Post(c *gin.Context) {
	username := c.PostForm("username") // 获取表单提交值
	pageSize := c.DefaultPostForm("pageSize", "10")
	c.JSON(http.StatusOK, gin.H{
		"name":     "张三",
		"username": username,
		"pageSize": pageSize,
	})
}
