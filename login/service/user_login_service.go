package service

import (
	"gin_demo/login/model"
	"gin_demo/login/serializer"
)

type UserLoginService struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,len=8"`
}

func (service *UserLoginService) Login() serializer.Response {
	sql := "select count(1) from user where username=? and password=?"
	var count int
	_ = model.DB.Get(&count, sql, service.UserName, service.Password)
	if count == 0 {
		return serializer.Response{
			Code: 40003,
			Msg:  "账号或密码错误",
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "登录成功",
	}
}
