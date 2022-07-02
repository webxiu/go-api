package main

import (
	"fmt"
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

	routers.AdminRouters(r)
	routers.ApiRouters(r)
	routers.DefaultRouters(r)

	r.Run() // r.Run(":8000") 默认端口:8080
}
