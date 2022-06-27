package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserList struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"` // 获取表单数据
}
type XMLType struct {
	UserName string `json:"username" xml:"username"`
	PassWord string `json:"password" xml:"password"` // 获取表单数据
}

func unixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()
	/** 必须反正r.LoadHTMLGlob之前 */
	r.SetFuncMap(template.FuncMap{
		"unixToTime": unixToTime,
	})
	/** 配置模板文件 */
	// r.LoadHTMLGlob("template/*") // 一层目录
	r.LoadHTMLGlob("template/**/*") // 二层目录

	/** 设置静态资源目录 */
	r.Static("/assets", "./static")

	r.GET("/", func(c *gin.Context) {
		page := c.Query("page")
		pageSize := c.DefaultQuery("pageSize", "10")
		c.JSON(http.StatusOK, gin.H{
			"name":     "张三",
			"page":     page,
			"pageSize": pageSize,
		})
	})
	// 1.获取参数
	r.POST("/post", func(c *gin.Context) {
		username := c.PostForm("username") // 获取表单提交值
		pageSize := c.DefaultPostForm("pageSize", "10")
		c.JSON(http.StatusOK, gin.H{
			"name":     "张三",
			"username": username,
			"pageSize": pageSize,
		})
	})
	// 2.GET或者POST绑定结构体返回数据, 写法一样
	r.GET("/get-struct", func(c *gin.Context) {
		userInfo := &UserList{}
		if err := c.ShouldBind(&userInfo); err == nil {
			c.JSON(http.StatusOK, userInfo)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}
	})
	// 3.解析XML数据, 使用 raw 提交xml 格式的数据
	r.GET("/get-xml", func(c *gin.Context) {
		xmlInfo := &XMLType{}
		xmlSliceData, _ := c.GetRawData() // 获取c.Request.Body中的数据
		fmt.Println(xmlSliceData)

		if err := xml.Unmarshal(xmlSliceData, &xmlInfo); err == nil {
			c.JSON(http.StatusOK, xmlInfo)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})
	// 动态路由
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(200, "%v", cid)
	})

	// 2.多文件夹r.LoadHTMLGlob("template/**/*")模板配置
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title": "首页数据",
			"time":  1656233903,
		})
	})
	r.GET("/home/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/goods.html", gin.H{
			"title": "首页数据goods",
		})
	})
	r.Run() // r.Run(":8000") 默认端口:8080
}
