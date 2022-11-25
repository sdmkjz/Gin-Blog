package v1

import (
	"gin-blog/pkg/e"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加分类
func AddCate(c *gin.Context) {
	var cateservice service.CateService
	_ = c.ShouldBindJSON(&cateservice)
	res := cateservice.AddCate()
	c.JSON(e.SUCCSE, res)
}

// 查询分类详情
func CateInfo(c *gin.Context) {
	var cateservice service.CateService
	id, _ := strconv.Atoi(c.Param("id"))
	res := cateservice.CateInfo(id)
	c.JSON(e.SUCCSE, res)
}

// 查询分类列表
func CateList(c *gin.Context) {
	var cateservice service.CateService
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	res := cateservice.CateList(pageSize, pageNum)
	c.JSON(e.SUCCSE, res)
}

// 删除分类
func DelCate(c *gin.Context) {
	var cateservice service.CateService
	id, _ := strconv.Atoi(c.Param("id"))
	res := cateservice.DelCate(id)
	c.JSON(e.SUCCSE, res)
}
