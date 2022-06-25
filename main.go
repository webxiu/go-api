package main

import (
	"net/http"
	"github.com/gin-gonic/gin")

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "张三",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "张三ping",
		})
	})
	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "这是返回的字符串")
	})
	// r.Run(":8000") 
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}