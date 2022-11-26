package v1

import (
	"gin-blog/pkg/e"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 创建文章
func AddArticle(c *gin.Context) {
	var art service.ArticleService
	_ = c.ShouldBindJSON(&art)
	res := art.AddArticle()
	c.JSON(e.SUCCSE, res)
}

// 查询单个文章
func ArticleInfo(c *gin.Context) {
	var art service.ArticleService
	id, _ := strconv.Atoi(c.Param("id"))
	res := art.ArticleInfo(id)
	c.JSON(e.SUCCSE, res)
}

// 修改单个文章
func ArticleEdit(c *gin.Context) {
	var art service.ArticleService
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&art)
	res := art.ArticleEdit(id)
	c.JSON(e.SUCCSE, res)
}

// 根据分类查询文章列表
func CateArticle(c *gin.Context) {
	var art service.ArticleService
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := art.CateArticle(id, pageSize, pageNum)
	c.JSON(e.SUCCSE, res)
}

// 删除文章
func ArticleDel(c *gin.Context) {
	var art service.ArticleService
	id, _ := strconv.Atoi(c.Param("id"))
	res := art.ArticleDel(id)
	c.JSON(e.SUCCSE, res)
}
