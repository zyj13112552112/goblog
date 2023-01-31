package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"
)

/*
	配置路由
*/

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default() //gin.Default()比gin.New()多加了两个中间间
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//配置路由组
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken()) //需要中间件
	{
		//User模块的路由
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由
		auth.POST("article/add", v1.AddArticle)      //添加文章
		auth.PUT("article/:id", v1.EditArticle)      //修改文章
		auth.DELETE("article/:id", v1.DeleteArticle) //删除文章
		//上传文件
		auth.POST("upload", v1.UpLoad)

	}
	rv1 := r.Group("/api/v1")
	{
		rv1.POST("user/add", v1.AddUser)
		rv1.GET("users", v1.GetUsers)
		rv1.GET("category", v1.GetCate)
		rv1.GET("article", v1.GetArticle)          //查询文章列表
		rv1.GET("article/list/:id", v1.GetCateArt) //查询分类下的所有文章
		rv1.GET("article/:id", v1.GetArtInfo)      //查询单个文章
		rv1.POST("login", v1.Login)                //登录
	}
	r.Run(utils.HttpPort)
}
