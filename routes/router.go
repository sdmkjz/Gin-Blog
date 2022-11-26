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
			// 用户操作
			authed.GET("user/info/:id", v1.UserInfo)
			authed.GET("users", v1.UsersInfo)
			authed.PUT("user/:id", v1.UserEdit)
			authed.DELETE("user/:id", v1.UserDelete)
			// 分类操作
			authed.POST("category/add", v1.AddCate)
			authed.GET("category/:id", v1.CateInfo)
			authed.GET("categorys", v1.CateList)
			authed.DELETE("category/:id", v1.DelCate)
			// 文章操作
			authed.POST("article/add", v1.AddArticle)
			authed.GET("article/:id", v1.ArticleInfo)
			authed.PUT("article/:id", v1.ArticleEdit)
			authed.GET("articles/:id", v1.CateArticle)
			authed.DELETE("article/:id", v1.ArticleDel)
		}
	}
	return r
}
