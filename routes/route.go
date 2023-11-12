package routes

import (
	common "gindiary/api/common"
	v1 "gindiary/api/v1"
	v2 "gindiary/api/xcx_api"

	// "log"
	// "os"

	"gindiary/model"
	"gindiary/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(util.AppMode)
	r := gin.Default()
	// 设置安全代理列表
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 定义需要进行 token 校验的中间件
	authMiddleware := model.AuthMiddleware()
	xcxMiddleware := model.XcxAuthMiddleware()
	tywMiddleware := model.TywAuthMiddleware()

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

		//天悦湾后台管理 社区区管理接口
		router.POST("/community/add", v2.AddCommunity)
		router.POST("/community/get", v2.GetCommunity)
		router.POST("/community/edit", v2.EdiCommunityState)
		router.POST("/community/det", v2.DeleteCommunity)

		//天悦湾后台管理 业主审核管理接口
		router.POST("/apply_list/add", v2.AddOwnerApply)
		router.POST("/apply_list/get", v2.GetOwnerApply)
		router.POST("/apply_list/edit", v2.EditApplyState)
		router.POST("/apply_list/det", v2.DeleteOwner)

		//天悦湾后台管理 帖子管理接口
		router.POST("/park_posts/get", v2.GetParkPosts)
		router.POST("/park_posts/edit", v2.EditPostsState)
		router.POST("/park_posts/det", v2.DeletePosts)

	}
	tyw_xcx_router := r.Group("api/v3")
	{

		//小程序登录注册
		tyw_xcx_router.POST("/user/login", v2.TywUserLogin)
		//token校验 --  以下接口都需要校验token，如果不想校验，请写在上边
		//用户模块
		tyw_xcx_router.Use(tywMiddleware)
		tyw_xcx_router.POST("/user/delete_user", v2.TywDeleteUser)
		tyw_xcx_router.POST("/user/check_user", v2.TywCheckUser)
		tyw_xcx_router.POST("/user/change_name", v2.TywChangeUserName)

		//七牛模块路由接口
		tyw_xcx_router.GET("/qiniu/token", common.GetQiNiuUpToken)
		tyw_xcx_router.POST("/qiniu/remove_file", common.DeleteQiNiuFile)

		//申请认证模块
		tyw_xcx_router.POST("/user/owner_apply", v2.AddOwnerApply)
		tyw_xcx_router.POST("/user/apply_list", v2.GetOwnerApply)
		tyw_xcx_router.POST("/user/check_apply", v2.EditApplyState)
		tyw_xcx_router.POST("/user/delete_apply", v2.DeleteOwner)

		//帖子模块
		tyw_xcx_router.POST("/user/add_park_posts", v2.AddParkPosts)
		tyw_xcx_router.POST("/user/park_posts_list", v2.GetParkPosts)
		tyw_xcx_router.POST("/user/edit_park_posts", v2.EditPostsState)
		tyw_xcx_router.POST("/user/delete_park_posts", v2.DeletePosts)

		//天悦湾后台管理 社区区管理接口
		tyw_xcx_router.POST("/community/add", v2.AddCommunity)
		tyw_xcx_router.POST("/community/get", v2.GetCommunity)
		tyw_xcx_router.POST("/community/edit", v2.EdiCommunityState)
		tyw_xcx_router.POST("/community/det", v2.DeleteCommunity)
	}
	xcx_router := r.Group("api/v2")
	{
		// //小程序登录注册
		// xcx_router.POST("/user/login", v2.XcxUserLogin)

		// xcx_router.POST("/user/free_analysis", v2.XcxFreeAnalysisURL)

		// //token校验 --  以下接口都需要校验token，如果不想校验，请写在上边
		// xcx_router.Use(xcxMiddleware)
		// // xcx_router.POST("/user/analysis", v2.XcxAnalysisURL)
		// xcx_router.POST("/user/analysis", v2.XcxOneSelfFreeAnalysisURL)

		// xcx_router.POST("/user/analysisRecord", v2.XcxGetAnalysis)
		// xcx_router.POST("/user/signIn", v2.XcxGetSignIn)
		//小程序登录注册
		xcx_router.POST("/user/login", v2.XcxUserLogin)

		xcx_router.POST("/user/free_analysis", v2.XcxFreeAnalysisURL)

		//token校验 --  以下接口都需要校验token，如果不想校验，请写在上边
		xcx_router.Use(xcxMiddleware)
		// xcx_router.POST("/user/analysis", v2.XcxAnalysisURL)
		xcx_router.POST("/user/analysis", v2.XcxOneSelfFreeAnalysisURL)

		xcx_router.POST("/user/analysisRecord", v2.XcxGetAnalysis)
		xcx_router.POST("/user/signIn", v2.XcxGetSignIn)
	}
	r.Run(util.HttpPort)

}
