package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}


func CreateSuccess(id int, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		200,
		id,
		"",
	})
}

func ParamError(message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		2,
		"",
		message,
	})
}

func DbError(message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		3,
		"",
		message,
	})
}

func LoginError(message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		4,
		"",
		message,
	})
}