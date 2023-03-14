package routes

import (
	common "gindiary/api/common"
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
		router.POST("/user/logout", v1.Logout)

		//七牛模块路由接口
		router.GET("/qiniu/token", common.GetQiNiuUpToken)
		router.POST("/qiniu/remove_file", common.DeleteQiNiuFile)
		//分类模块路由接口
		router.POST("/category/add", v1.AddCategory)
		router.POST("/category/get", v1.GetCategory)
		router.POST("/category/edit", v1.EditCategory)
		router.POST("/category/det", v1.DeleteCategory)
		router.POST("/category/detail", v1.GetCategoryDet)

		//文章模块路由接口
		router.POST("/article/add", v1.AddArticle)
		router.POST("/article/get", v1.GetArticleInfo)
		router.POST("/article/edit", v1.EditArticle)
		router.POST("/article/det", v1.DeleteArticle)
		router.POST("/article/catelist", v1.GetCateArts)
		router.POST("/article/artlist", v1.GetArts)

	}
	r.Run(util.HttpPort)

}
