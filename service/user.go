package service

import (
	"errors"
	"go-demo/global"
	"go-demo/model"
	"go-demo/model/request"
	"go-demo/utils"
	"time"
)

func LoginByPassword(login request.LoginStruct) (user model.User, err error) {
	login.Password = utils.MD5([]byte(login.Password))
	user.Phone = login.Phone
	user.Password = login.Password
	err = global.GVA_DB.Where(&user).First(&user).Error
	return
}

func Register(register request.RegisterStruct) (user model.User, err error){
	user.Avatar = register.Avatar
	user.Password = register.Password
	user.Phone = register.Phone
	user.Username = register.Username
	user.CreateAt = time.Now()
	//判断手机号是否已注册
	global.GVA_DB.Where("phone = ?", register.Phone).First(&user)
	if user.Id != 0 {
		return user , errors.New("手机号已被注册")
	}
	err = global.GVA_DB.Create(&user).Error
	return
}