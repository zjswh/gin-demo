package initialize

import (
	"github.com/gin-gonic/gin"
	"go-demo/global/response"
	"go-demo/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//处理跨域
	Router.Use(cors())

	ApiGroup := Router.Group("v1/api/")
	router.InitUserRouter(ApiGroup) //注册用户相关接口路由
	router.InitBaseRouter(ApiGroup) //注册用户相关接口路由

	//处理404
	Router.NoMethod(HandleNotFind)
	Router.NoRoute(HandleNotFind)

	return Router
}

//处理404
func HandleNotFind(c *gin.Context)  {
	response.Result(4, "api不存在", "", c)
}

//跨域
func cors() gin.HandlerFunc  {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//过滤options请求
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}