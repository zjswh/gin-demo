package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-demo/api/v1"
	"go-demo/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").
		Use(middleware.JwtAuth())
	{
		userRouter.GET("getInfo", v1.GetInfo)

	}
}
