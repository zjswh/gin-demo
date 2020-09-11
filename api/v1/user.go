package v1

import (
	"github.com/gin-gonic/gin"
	"go-demo/global/response"
	"go-demo/middleware"
	"go-demo/model/request"
	"go-demo/service"
	"go-demo/utils"
)

func Login(c *gin.Context) {
	//获取参数
	var param request.LoginStruct
	_ = c.ShouldBindJSON(&param)

	UserVerify := utils.Rules{
		"Phone": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(param, UserVerify)
	if UserVerifyErr != nil {
		response.ParamError(UserVerifyErr.Error(), c)
		return
	}
	user, err := service.LoginByPassword(param)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	//生成token
	var UserClaims request.UserClaims
	UserClaims.Id = user.Id
	UserClaims.Username = user.Username
	UserClaims.Phone = user.Phone
	UserClaims.Avatar = user.Avatar
	token, err := middleware.CreateToken(UserClaims)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Result(200, token, "", c)
	return
}

func Register(c *gin.Context) {
	var param request.RegisterStruct
	_ = c.ShouldBindJSON(&param)
	RegisterVerify := utils.Rules{
		"Phone": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
		"Avatar": {utils.NotEmpty()},
		"Username": {utils.NotEmpty()},
	}

	RegisterVerifyErr := utils.Verify(param, RegisterVerify)
	if RegisterVerifyErr != nil {
		response.ParamError(RegisterVerifyErr.Error(), c)
		return
	}
	user, err := service.Register(param)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}

	response.CreateSuccess(user.Id, c)
	return
}

func GetInfo(c *gin.Context)  {
	response.Result(200, "asd", "", c)
	return
}
