package api

import (
	"fmt"
	"gin-api/controllers/base"
	"gin-api/utils"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	base.BaseGlobalController
}

// 单个上传
func (con UploadController) Single(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("file")
	savePath := path.Join("./static/upload", file.Filename)
	url := path.Join(c.Request.Host, "./assets/upload", file.Filename)

	if err == nil {
		c.SaveUploadedFile(file, savePath)
	}

	con.Success(c, "上传成功", gin.H{
		"username": username,
		"url":      url,
	})
}

// 批量上传
func (con UploadController) Batch(c *gin.Context) {
	username := c.PostForm("username")
	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	urls := []string{} // var urls []string
	// 循环保存并获取路径
	for _, file := range files {
		savePath := path.Join("./static/upload", file.Filename)
		visitPath := path.Join(c.Request.Host, "./assets/upload", file.Filename)
		c.SaveUploadedFile(file, savePath)
		urls = append(urls, visitPath)
	}

	con.Success(c, "上传成功", gin.H{
		"username": username,
		"urls":     urls,
	})
}

// 按日期目录保存
func (con UploadController) DirSave(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("file")

	if err == nil {
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		if _, ok := allowExtMap[extName]; !ok {
			c.String(200, "文件类型不合法")
			return
		}

		day := utils.GetDay()
		dir := "./static/upload/" + day
		visitPath := "./assets/upload/" + day
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Println(err)
			c.String(200, "目录创建失败")
			return
		}
		fileName := strconv.FormatInt(utils.GetUnix(), 10) + extName
		savePath := path.Join(dir, fileName)
		c.SaveUploadedFile(file, savePath)

		url := path.Join(c.Request.Host, visitPath, fileName)
		con.Success(c, "上传成功", gin.H{
			"username": username,
			"url":      url,
		})
	}

}
