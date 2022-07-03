package main

import (
	"fmt"
	"gin-api/middlewares"
	routers "gin-api/router"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func unixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default() // gin.New() 创建一个新的路由, 不使用默认的日制等配置
	/** 必须放在r.LoadHTMLGlob之前 */
	r.SetFuncMap(template.FuncMap{
		"unixToTime": unixToTime,
	})
	/** 配置模板文件 */
	// r.LoadHTMLGlob("template/*") // 一层目录
	r.LoadHTMLGlob("template/**/*") // 二层目录

	/** 设置静态资源目录 */
	r.Static("/assets", "./static")

	// 全局中间件
	r.Use(middlewares.Middleware{}.Middleware1, middlewares.Middleware{}.Middleware2)

	routers.ApiRouters(r)
	routers.AdminRouters(r)
	routers.HomeRouters(r)

	r.Run() // r.Run(":8000") 默认端口:8080
}
