package v1

import (
	"fmt"
	"gin-blog/pkg/e"
	"gin-blog/pkg/utils"
	"gin-blog/serializer"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
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
