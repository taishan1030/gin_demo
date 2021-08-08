package service

import (
	"gin_demo/login/model"
	"gin_demo/login/serializer"
)

type UserRegisterSerivce struct {
	UserName string `json:"username" binding:"required,min=3,max=7" db:"username"`
	Password string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (service *UserRegisterSerivce) Register() serializer.Response {
	sql := `select count(1) from user where email = ?`
	var count int
	_ = model.DB.Get(&count, sql, service.Email)
	if count > 0 {
		return serializer.Response{
			Code:  40001,
			Data:  nil,
			Msg:   "邮箱已经被注册",
			Error: "",
		}
	}

	sqlStr := "insert into user(username,password,age,sex,email) values(:username,:password,:age,:sex,:email)"
	if _, err := model.DB.NamedExec(sqlStr, service); err != nil {
		return serializer.Response{
			Code:  40002,
			Data:  nil,
			Msg:   "注册失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "注册成功",
		Error: "",
	}

}
