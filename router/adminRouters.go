package routers

import (
	"gin-api/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/index", admin.UserController{}.Index)
		// 动态路由
		adminRouter.GET("/list/:cid", admin.UserController{}.List)
		// 解析XML数据, 使用 raw 提交xml 格式的数据
		adminRouter.GET("/get-xml", admin.UserController{}.GetXml)
		// GET或者POST绑定结构体返回数据, 写法一样
		adminRouter.GET("/get-struct", admin.UserController{}.GetStruct)
		// test success
		adminRouter.GET("/test-success", admin.UserController{}.TestSuccess)
		// test error
		adminRouter.GET("/test-error", admin.UserController{}.TestError)
	}
}
