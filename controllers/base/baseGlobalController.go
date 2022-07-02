package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseGlobalController struct{}

// 成功
func (con BaseGlobalController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "成功-全局",
		"data":   data,
	})
}

// 失败
func (con BaseGlobalController) Error(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "失败-全局",
		"data":   data,
	})
}
