package routes

import (
	v1 "gin-blog/api/v1"
	"gin-blog/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("user/register", v1.UserRegister)
		apiv1.POST("user/login", v1.UserLogin)
		authed := apiv1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("user/info/:id", v1.UserInfo)
			authed.GET("users", v1.UsersInfo)
			authed.PUT("user/:id", v1.UserEdit)
			authed.DELETE("user/:id", v1.UserDelete)
		}
	}
	return r
}
