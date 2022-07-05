package api

import (
	"gin-api/controllers/base"
	"path"

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

	con.Success(c, gin.H{
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

	con.Success(c, gin.H{
		"username": username,
		"urls":     urls,
	})
}
