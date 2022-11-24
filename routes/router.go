package routes

import (
	v1 "gin-blog/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("user/register", v1.UserRegister)
		apiv1.POST("user/login", v1.UserLogin)
	}
	return r
}
