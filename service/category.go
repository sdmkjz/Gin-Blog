package service

import (
	"gin-blog/model"
	"gin-blog/pkg/e"
	"gin-blog/serializer"
	"gorm.io/gorm"
)

type CateService struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (service *CateService) AddCate() serializer.Response {
	var cate model.Category
	var count int64
	model.Db.Where("name = ?", service.Name).First(&cate).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: e.ERROR_CATENAME_USED,
			Msg:    e.GetErrMsg(e.ERROR_CATENAME_USED),
		}
	}
	cate.Name = service.Name
	if err := model.Db.Create(&cate).Error; err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    e.GetErrMsg(e.ERROR),
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}

func (service *CateService) CateInfo(id int) serializer.Response {
	var cate model.Category
	if err := model.Db.Where("Id = ?", id).First(&cate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR_CATE_NOT_EXIST,
				Msg:    e.GetErrMsg(e.ERROR_CATE_NOT_EXIST),
			}
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Data:   cate,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}

func (service *CateService) CateList(pageSize int, pageNum int) serializer.ListResponse {
	var cate []model.Category
	var total int64
	model.Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total)
	return serializer.ListResponse{
		Status: e.SUCCSE,
		Data:   cate,
		Total:  total,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}

func (service *CateService) DelCate(id int) serializer.Response {
	var cate model.Category
	if err := model.Db.Where("id = ?", id).First(&cate).Error; err != nil {
		return serializer.Response{
			Status: e.ERROR_CATE_NOT_EXIST,
			Msg:    e.GetErrMsg(e.ERROR_CATE_NOT_EXIST),
		}
	}
	model.Db.Where("id = ?", id).Delete(&cate)
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}
