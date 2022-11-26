package service

import (
	"gin-blog/model"
	"gin-blog/pkg/e"
	"gin-blog/serializer"
	"gorm.io/gorm"
	"log"
)

type ArticleService struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Cid     uint   ` json:"cid"`
	Content string `json:"content"`
	Img     string ` json:"img"`
}

func (service *ArticleService) AddArticle() serializer.Response {
	var art model.Article
	art.Title = service.Title
	art.Desc = service.Desc
	art.Cid = service.Cid
	art.Content = service.Content
	art.Img = service.Img
	if err := model.Db.Create(&art).Error; err != nil {
		log.Println(err)
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

func (service *ArticleService) ArticleInfo(id int) serializer.Response {
	var art model.Article
	if err := model.Db.Preload("Category").Where("id = ?", id).First(&art).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR_ART_NOT_EXIST,
				Msg:    e.GetErrMsg(e.ERROR_ART_NOT_EXIST),
			}
		}
	}
	return serializer.Response{
		Status: e.SUCCSE,
		Data:   art,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}

func (service *ArticleService) ArticleEdit(id int) serializer.Response {
	var art model.Article
	art.Title = service.Title
	art.Desc = service.Desc
	art.Cid = service.Cid
	art.Content = service.Content
	art.Img = service.Img
	if err := model.Db.Where("id = ?", id).First(&art).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR_ART_NOT_EXIST,
				Msg:    e.GetErrMsg(e.ERROR_ART_NOT_EXIST),
			}
		}

	}
	model.Db.Where("id = ?", id).Updates(&art)
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}

func (service *ArticleService) CateArticle(id int, pageSize int, pageNum int) serializer.ListResponse {
	var artlist []model.Article
	var total int64
	model.Db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", id).Find(&artlist).Count(&total)
	return serializer.ListResponse{
		Status: e.SUCCSE,
		Data:   artlist,
		Msg:    e.GetErrMsg(e.SUCCSE),
		Total:  total,
	}
}

func (service *ArticleService) ArticleDel(id int) serializer.Response {
	var art model.Article
	if err := model.Db.Where("id = ?", id).First(&art).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: e.ERROR_ART_NOT_EXIST,
				Msg:    e.GetErrMsg(e.ERROR_ART_NOT_EXIST),
			}
		}
	}
	model.Db.Where("id = ?", id).Delete(&art)
	return serializer.Response{
		Status: e.SUCCSE,
		Msg:    e.GetErrMsg(e.SUCCSE),
	}
}
