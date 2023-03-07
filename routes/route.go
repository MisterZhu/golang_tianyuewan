package routes

import (
	v1 "gindiary/api/v1"
	"gindiary/model"
	"gindiary/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(util.AppMode)
	r := gin.Default()

	// 定义需要进行 token 校验的中间件
	authMiddleware := model.AuthMiddleware()

	router := r.Group("api/v1")
	{

		//用户模块路由接口
		router.POST("/user/register", v1.Register)
		router.POST("/user/login", v1.Login)

		//token校验 --  以下接口都需要校验token，如果不想校验，请写在上边
		router.Use(authMiddleware)

		router.POST("/user/update", v1.EditUser)
		router.GET("/user/info", v1.Info)

		//分类模块路由接口
		router.POST("/category/add", v1.AddCategory)
		router.POST("/category/get", v1.GetCategory)
		router.POST("/category/edit", v1.EditCategory)
		router.POST("/category/det", v1.DeleteCategory)

		//文章模块路由接口
		router.POST("/article/add", v1.AddArticle)
		router.POST("/article/get", v1.GetArticle)
		router.POST("/article/edit", v1.EditArticle)
		router.POST("/article/det", v1.DeleteArticle)
	}
	r.Run(util.HttpPort)

}
