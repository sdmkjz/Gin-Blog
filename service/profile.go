package service

import (
	"gin-blog/model"
	"gin-blog/pkg/e"
	"gin-blog/serializer"
	"gorm.io/gorm"
)

type ProfileService struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	QQchat    string `json:"qq_chat"`
	Wechat    string `json:"wechat"`
	Weibo     string `json:"weibo"`
	Bili      string `json:"bili"`
	Email     string `json:"email"`
	Img       string `json:"img"`
	Avatar    string `json:"avatar"`
	IcpRecord string `json:"icp_record"`
}

func (service *ProfileService) ProfileUpdate(id int) serializer.Response {
	var prof model.Profile
	//prof.Name = service.Name
	//prof.Desc = service.Desc
	//prof.QQchat = service.QQchat
	//prof.Wechat = service.Wechat
	//prof.Weibo = service.Weibo
	//prof.Bili = service.Bili
	//prof.Email = service.Email
	//prof.Img = service.Img
	//prof.Avatar = service.Avatar
	//prof.IcpRecord = service.IcpRecord
	if err := model.Db.Where("id = ?", id).First(&prof).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			model.Db.Model(&prof).Where("id = ?", id).Create(&service)
			return serializer.Response{
				Status: e.SUCCSE,
				Msg:    e.GetErrMsg(e.SUCCSE),
			}
		}
	}
	model.Db.Model(&prof).Where("id = ?", id).Updates(&service)
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}

func (service *ProfileService) ProfileInfo(id int) serializer.Response {
	var prof model.Profile
	if err := model.Db.Where("id = ?", id).First(&prof).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR,
				Msg:    e.GetErrMsg(e.ERROR),
			}
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Data:   prof,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}
