package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// 成功
func (con BaseController) success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "成功-单个包",
		"data":   data,
	})
}

// 失败
func (con BaseController) error(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "失败-单个包",
		"data":   data,
	})
}
