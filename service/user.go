package service

import (
	"fmt"
	"gin-blog/model"
	"gin-blog/serializer"
)

type UserService struct {
	Username string `form:"username" json:"username" validate:"required,min=4,max=12" label:"用户名""`
	Password string `form:"password" json:"password" validate:"required,min=6,max=20" label:"密码"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.Db.Model(&user).Where("username=?", service.Username).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户名已存在",
		}
	}
	user.Username = service.Username
	// 密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.Db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return serializer.Response{
			Status: 400,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}
}
