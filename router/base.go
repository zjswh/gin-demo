package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-demo/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("base")
	{
		userRouter.POST("login", v1.Login)
		userRouter.POST("register", v1.Register)
	}
}
