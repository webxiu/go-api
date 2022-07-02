package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (con HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "首页数据",
		"time":  1656233903,
	})
}
func (con HomeController) Page(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.DefaultQuery("pageSize", "10")
	c.JSON(http.StatusOK, gin.H{
		"name":     "张三",
		"page":     page,
		"pageSize": pageSize,
	})
}
func (con HomeController) Goods(c *gin.Context) {
	c.HTML(http.StatusOK, "home/goods.html", gin.H{
		"title": "首页数据goods",
	})
}
