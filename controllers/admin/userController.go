package admin

import (
	"encoding/xml"
	"fmt"
	"gin-api/controllers/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// 单个包继承BaseController
	BaseController
	// 全局继承BaseController2
	base.BaseGlobalController
}

type XMLType struct {
	UserName string `json:"username" xml:"username"`
	PassWord string `json:"password" xml:"password"` // 获取表单数据
}

type UserList struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"` // 获取表单数据
}

func (con UserController) Index(c *gin.Context) {
	cid := c.Param("cid")
	// c.String(200, "%v", cid)
	// con.success(c, cid)
	con.Success(c, "成功了", cid)
}
func (con UserController) List(c *gin.Context) {
	cid := c.Param("cid")
	username, _ := c.Get("username")
	v, ok := username.(string)
	if ok {
		fmt.Printf("list, 中间件数据为:" + v + "\n")
	} else {
		fmt.Printf("list, 失败")
	}
	// c.String(200, "%v", cid)
	con.Error(c, 400, "参数错误"+cid)
}
func (con UserController) TestSuccess(c *gin.Context) {
	cid := c.Param("cid")
	con.success(c, "成功了"+cid)
}
func (con UserController) TestError(c *gin.Context) {
	cid := c.Param("cid")
	con.error(c, cid)
}

func (con UserController) GetXml(c *gin.Context) {
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
}
func (con UserController) GetStruct(c *gin.Context) {
	userInfo := &UserList{}
	if err := c.ShouldBind(&userInfo); err == nil {
		c.JSON(http.StatusOK, userInfo)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
}
