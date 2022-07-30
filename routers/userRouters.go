package routers

import (
	"gin-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("/list", user.UserController{}.List)
		userRouter.POST("/add", user.UserController{}.Add)
		userRouter.POST("/delete", user.UserController{}.Delete)
		userRouter.PUT("/update/:id", user.UserController{}.Update)
	}
}
