package v1

import (
	"fmt"
	"gin-blog/pkg/e"
	"gin-blog/pkg/utils"
	"gin-blog/serializer"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 用户注册
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	_ = c.ShouldBindJSON(&userRegister)
	if msg, code := utils.Validate(&userRegister); code == e.SUCCSE {
		res := userRegister.Register()
		c.JSON(e.SUCCSE, res)
	} else {
		fmt.Println(code)
		c.JSON(e.ERROR, serializer.Response{
			Status: code,
			Msg:    msg,
		})
	}
}

// 用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	_ = c.ShouldBindJSON(&userLogin)
	if msg, code := utils.Validate(&userLogin); code == e.SUCCSE {
		res := userLogin.Login()
		c.JSON(e.SUCCSE, res)
	} else {
		fmt.Println(code)
		c.JSON(e.ERROR, serializer.Response{
			Status: code,
			Msg:    msg,
		})
	}
}

// 查询单个用户
func UserInfo(c *gin.Context) {
	var userinfo service.UserService
	id, _ := strconv.Atoi(c.Param("id"))
	data := userinfo.Info(id)
	c.JSON(e.SUCCSE, data)
}

// 查询用户列表
func UsersInfo(c *gin.Context) {
	var usersinfo service.UserService
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := usersinfo.InfoList(username, pageSize, pageNum)
	c.JSON(e.SUCCSE, data)
}
