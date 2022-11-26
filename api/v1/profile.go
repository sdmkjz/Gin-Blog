package v1

import (
	"gin-blog/pkg/e"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 创建/更新个人信息
func ProfilePut(c *gin.Context) {
	var prof service.ProfileService
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&prof)
	res := prof.ProfileUpdate(id)
	c.JSON(e.SUCCSE, res)
}

// 获取个人信息
func ProfileInfo(c *gin.Context) {
	var prof service.ProfileService
	id, _ := strconv.Atoi(c.Param("id"))
	res := prof.ProfileInfo(id)
	c.JSON(e.SUCCSE, res)
}
