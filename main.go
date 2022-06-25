package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserList struct {
	UserName string `json:"username"`
	PassWord int    `json:"password"`
}

func main() {
	r := gin.Default()
	// 配置模板文件
	// r.LoadHTMLGlob("template/*") // 一层目录
	r.LoadHTMLGlob("template/**/*") // 二层目录

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "张三",
		})
	})
	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "这是返回的字符串")
	})
	r.GET("/struct", func(c *gin.Context) {
		result := &UserList{
			UserName: "李四",
			PassWord: 44,
		}
		c.JSON(200, result)
	})
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"name": "张三ping",
			"age":  19,
		})
	})
	r.GET("/jsonp", func(c *gin.Context) {
		// http://localhost:8080/jsonp?callback=cb
		c.JSONP(200, map[string]interface{}{
			"name": "张三-jsonp",
			"age":  19,
		})
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "返回 xml 数据",
		})
	})

	// 1.只能使用r.LoadHTMLGlob("template/*")的模板配置
	r.GET("/tpl-news", func(c *gin.Context) {
		info := &UserList{
			UserName: "李四",
			PassWord: 44,
		}
		c.HTML(http.StatusOK, "/news.html", gin.H{
			"title": "模板数据news",
			"info":  info,
		})
	})
	// 只能使用r.LoadHTMLGlob("template/*")的模板配置
	r.GET("/tpl-goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/goods.html", gin.H{
			"title": "模板数据goods",
		})
	})

	// 2.多文件夹r.LoadHTMLGlob("template/**/*")模板配置
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "首页数据",
		})
	})
	r.GET("/home/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/goods.html", gin.H{
			"title": "首页数据goods",
		})
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台数据index",
		})
	})
	r.GET("/admin/news", func(c *gin.Context) {
		info := &UserList{
			UserName: "新闻页面",
			PassWord: 88,
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "后台数据news",
			"info":  info,
			"hobby": []string{"吃饭", "睡觉", "打豆豆"},
			"newsList": []interface{}{
				&UserList{
					UserName: "新闻页面111",
					PassWord: 1111,
				}, &UserList{
					UserName: "新闻页面222",
					PassWord: 2222,
				}, &UserList{
					UserName: "新闻页面333",
					PassWord: 3333,
				},
			},
		})
	})
	r.Run() // r.Run(":8000") 默认端口:8080
}
