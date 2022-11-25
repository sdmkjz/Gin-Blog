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
