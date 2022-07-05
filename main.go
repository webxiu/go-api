package main

import (
	"gin-api/middlewares"
	"gin-api/routers"
	"gin-api/utils"

	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // gin.New() 创建一个新的路由, 不使用默认的日制等配置
	/** 必须放在r.LoadHTMLGlob之前 */
	r.SetFuncMap(template.FuncMap{"unixToTime": utils.UnixToTime})
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
	routers.UploadRouters(r)

	r.Run() // r.Run(":8000") 默认端口:8080
}
