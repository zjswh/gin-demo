package model

import "time"

type User struct {
	Id	int `json:"id" gorm:"comment:'唯一id'"`
	Phone string `json:"phone" gorm:"comment:'用户手机号'"`
	Password string `json:"-" gorm:"comment:'用户密码'"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
}