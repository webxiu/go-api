package user

import (
	"fmt"
	"gin-api/controllers/base"
	"gin-api/mysql"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseGlobalController
}

var mDB = mysql.Connect()

// 查询列表
func (con UserController) List(c *gin.Context) {
	var dataList []mysql.UserList
	var total int64
	username := c.Query("username")
	phone := c.Query("phone")
	status, _ := strconv.Atoi(c.Query("status"))
	address := c.Query("address")
	email := c.Query("email")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	fmt.Println("======PAGE:", page, pageSize)

	query := &mysql.UserList{UserName: username, Phone: phone, Status: status, Address: address, Email: email}

	mDB.Model(dataList).Count(&total).Limit(pageSize).Offset(page).Where(query).Find(&dataList)

	fmt.Println("total:", total)
	if len(dataList) == 0 {
		con.Result(c, 200, "没有查询到数据")
	} else {
		con.Result(c, 200, "查询成功", gin.H{
			"list":  dataList,
			"total": total,
		})
	}
}

// 添加用户
func (con UserController) Add(c *gin.Context) {
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	status, _ := strconv.Atoi(c.PostForm("status"))
	address := c.PostForm("address")
	email := c.PostForm("email")
	query := mysql.UserList{UserName: username, Phone: phone, Status: status, Address: address, Email: email}
	/** 批量添加 */
	// querys := []mysql.UserList{
	// 	{UserName: username, Phone: phone, Status: status, Address: address, Email: email},
	// 	{UserName: username, Phone: phone, Status: status, Address: address, Email: email},
	// 	{UserName: username, Phone: phone, Status: status, Address: address, Email: email},
	// }
	// result := mDB.CreateInBatches(&querys, 100)

	result := mDB.Create(&query) // 通过数据的指针来创建
	fmt.Println("行", result.RowsAffected)

	if result.RowsAffected > 0 {
		con.Result(c, 200, "添加成功")
	} else {
		con.Result(c, 400, "添加失败")
	}

}

// 根据id删除用户
func (con UserController) Delete(c *gin.Context) {
	var data []mysql.UserList
	id := c.PostForm("id")
	mDB.Where("id = ?", id).Find(&data)
	if len(data) == 0 {
		con.Result(c, 400, "没有找到删除的内容")
		return
	}
	mDB.Where("id = ?", id).Delete(&data)
	con.Result(c, 200, "删除成功")

}

// 根据id修改用户
func (con UserController) Update(c *gin.Context) {
	var data mysql.UserList
	// id := c.Param("id")
	id, _ := strconv.Atoi(c.Param("id"))
	// id, _ := strconv.Atoi(c.PostForm("id"))
	// username := c.PostForm("username")
	// phone := c.PostForm("phone")
	// status, _ := strconv.Atoi(c.PostForm("status"))
	// address := c.PostForm("address")
	// email := c.PostForm("email")
	// query := mysql.UserList{UserName: username, Phone: phone, Status: status, Address: address, Email: email}

	fmt.Println("====id:", id)
	mDB.Select("id").Where("id = ?", id).Find(&data)
	// id 是否存在
	if data.ID == 0 {
		con.Result(c, 400, "没有找到修改的ID")
		return
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		con.Result(c, 400, "修改失败")
	} else {
		mDB.Where("id = ?", id).Updates(&data)
		con.Result(c, 200, "修改成功", data)
	}
}
