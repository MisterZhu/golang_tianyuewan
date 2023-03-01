package routes

import (
	v1 "gindiary/api/v1"
	"gindiary/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(util.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{

		//用户模块路由接口
		router.POST("/user/register", v1.Register)
		router.POST("/user/login", v1.Login)
		router.POST("/user/update", v1.EditUser)

		// router.PUT("/:id", v1.EditUser)
		// router.DELETE("/:id", v1.DeleteUser)
		// router.POST("/login", api.Login)
		// router.GET("/info", middlewares.AuthMiddleware(), api.Info)

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
