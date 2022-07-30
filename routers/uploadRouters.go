package routers

import (
	"gin-api/controllers/api"

	"github.com/gin-gonic/gin"
)

func UploadRouters(r *gin.Engine) {
	apiRouter := r.Group("/upload")
	{
		apiRouter.POST("/single", api.UploadController{}.Single)
		apiRouter.POST("/batch", api.UploadController{}.Batch)
		apiRouter.POST("/dir-save", api.UploadController{}.DirSave)
	}
}
