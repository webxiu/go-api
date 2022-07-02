package routers

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type XMLType struct {
	UserName string `json:"username" xml:"username"`
	PassWord string `json:"password" xml:"password"` // 获取表单数据
}

type UserList struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"` // 获取表单数据
}

func AdminRouters(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		// 动态路由
		adminRouter.GET("/list/:cid", func(c *gin.Context) {
			cid := c.Param("cid")
			c.String(200, "%v", cid)
		})
		// 解析XML数据, 使用 raw 提交xml 格式的数据
		adminRouter.GET("/get-xml", func(c *gin.Context) {
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

		// 2.GET或者POST绑定结构体返回数据, 写法一样
		adminRouter.GET("/get-struct", func(c *gin.Context) {
			userInfo := &UserList{}
			if err := c.ShouldBind(&userInfo); err == nil {
				c.JSON(http.StatusOK, userInfo)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		})
	}
}
