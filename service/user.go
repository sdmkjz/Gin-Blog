package service

import (
	"fmt"
	"gin-blog/model"
	"gin-blog/pkg/e"
	"gin-blog/pkg/utils"
	"gin-blog/serializer"
	"gorm.io/gorm"
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

func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.Db.Where("username=?", service.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR,
				Msg:    e.GetErrMsg(e.ERROR_USER_NOT_EXIST),
			}
		}
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR),
		}
	}
	// 验证密码
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR_PASSWORD_WRONG),
		}
	}
	// token返回
	token, err := utils.GenerateToken(user.ID, service.Username)
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR_TOKEN_WRONG),
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Data: serializer.TokenData{
			User:  user.Username,
			Token: token,
		},
	}
}
