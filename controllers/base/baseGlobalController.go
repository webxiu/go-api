package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseGlobalController struct{}

// 全局成功
func (con BaseGlobalController) Success(c *gin.Context, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    msg,
		"data":   data,
	})
}

// 全局失败
func (con BaseGlobalController) Error(c *gin.Context, status int, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}

// 全局成功与失败
func (con BaseGlobalController) Result(c *gin.Context, status int, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}
