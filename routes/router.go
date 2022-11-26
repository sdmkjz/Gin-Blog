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
		apiv1.GET("article/:id", v1.ArticleInfo)
		apiv1.GET("articles/:id", v1.CateArticle)
		apiv1.GET("category/:id", v1.CateInfo)
		apiv1.GET("categorys", v1.CateList)
		apiv1.GET("profile/:id", v1.ProfileInfo)
		authed := apiv1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.GET("user/info/:id", v1.UserInfo)
			authed.GET("users", v1.UsersInfo)
			authed.PUT("user/:id", v1.UserEdit)
			authed.DELETE("user/:id", v1.UserDelete)
			// 分类操作
			authed.POST("category/add", v1.AddCate)
			authed.DELETE("category/:id", v1.DelCate)
			// 文章操作
			authed.POST("article/add", v1.AddArticle)
			authed.PUT("article/:id", v1.ArticleEdit)
			authed.DELETE("article/:id", v1.ArticleDel)
			// 个人信息
			authed.PUT("profile/:id", v1.ProfilePut)
		}
	}
	return r
}
