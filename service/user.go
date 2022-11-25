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
	Role     int    `gorm:"type:int" json:"role" label:"角色"` // 1 管理员 2 普通用户
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.Db.Model(&user).Where("username=?", service.Username).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR_USERNAME_USED),
		}
	}
	user.Username = service.Username
	// 密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR),
		}
	}
	// 创建用户
	if err := model.Db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR),
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.REGISTER_SUCCSE),
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
	// 权限验证
	if user.Role != 1 {
		return serializer.Response{
			Status: e.ERROR_USER_NO_RIGHT,
			Msg:    e.GetErrMsg(e.ERROR_USER_NO_RIGHT),
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

func (service *UserService) Info(id int) serializer.Response {
	var user model.User
	if err := model.Db.Where("id=?", id).First(&user).Error; err != nil {
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
	return serializer.Response{
		Status: e.SUCCSE,
		Data:   user,
		Msg:    "ok",
	}
}

func (service *UserService) InfoList(username string, pageSize int, pageNum int) serializer.ListResponse {
	var user []model.User
	var total int64
	if username == "" {
		model.Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user).Count(&total)
		//model.Db.Model(&user).Count(&total)
		return serializer.ListResponse{
			Status: e.SUCCSE,
			Data:   user,
			Total:  total,
			Msg:    "ok",
		}
	}
	err := model.Db.Where("username LIKE ?", username+"%").Find(&user).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		return serializer.ListResponse{Status: e.ERROR, Error: e.GetErrMsg(e.ERROR)}
	}
	return serializer.ListResponse{
		Status: e.SUCCSE,
		Data:   user,
		Total:  total,
		Msg:    "ok",
	}

}

func (service *UserService) Edit(id int) serializer.Response {
	var user model.User
	var maps = make(map[string]interface{})
	if err := model.Db.Where("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR,
				Msg:    e.GetErrMsg(e.ERROR_USER_NOT_EXIST),
			}
		}
	}
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR),
		}
	}
	maps["username"] = service.Username
	maps["password"] = user.Password
	maps["role"] = service.Role
	if err := model.Db.Model(&user).Where("id = ?", id).Updates(maps).Error; err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR),
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Data:   user,
		Msg:    "ok",
	}
}

func (service *UserService) Delete(id int) serializer.Response {
	var user model.User
	if err := model.Db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR_USER_NOT_EXIST,
				Msg:    e.GetErrMsg(e.ERROR_USER_NOT_EXIST),
			}
		}
	}
	model.Db.Where("id = ?", id).Delete(&user)
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}
