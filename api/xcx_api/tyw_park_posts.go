package xcxapi

import (
	"fmt"
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"
	"log"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
发布帖子
*/
func AddParkPosts(c *gin.Context) {
	var posts model.TywParkPostsModel
	if err := c.ShouldBindJSON(&posts); err != nil {
		// 处理参数绑定错误
		log.Printf("Error binding request data: %v", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}

	// 可以在这里对 posts 进行一些其他处理，然后插入数据库等操作
	log.Printf("posts = %+v", posts)
	model.TywCreateParkPosts(&posts)
	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
}

// func AddParkPosts(c *gin.Context) {
// 	var posts = model.TywParkPostsModel{}
// 	c.ShouldBindJSON(&posts)

// 	var requestData struct {
// 		UserId        string `form:"user_id"`
// 		WeiXin        string `form:"weixin"`
// 		PostsType     int    `form:"posts_type"`
// 		ImgUrl        string `form:"img_url"`
// 		Telephone     string `form:"telephone"`
// 		InMaintenance bool   `form:"in_maintenance"`
// 		Negotiable    bool   `form:"negotiable"`
// 		Title         string `form:"title"`
// 	}

// 	if err := c.ShouldBind(&requestData); err != nil {
// 		// 处理参数绑定错误
// 		log.Printf("Error binding request data: %v", err)
// 		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
// 		return
// 	}

// 	data := model.TywParkPostsModel{
// 		UserId:        requestData.UserId,
// 		WeiXin:        requestData.WeiXin,
// 		State:         0,
// 		PostsType:     requestData.PostsType,
// 		ImgUrl:        requestData.ImgUrl,
// 		Telephone:     requestData.Telephone,
// 		InMaintenance: requestData.InMaintenance,
// 		Negotiable:    requestData.Negotiable,
// 		Title:         requestData.Title,
// 	}
// 	log.Printf("data = %+v", data)
// 	log.Printf("c = %v", c)
// 	model.TywCreateParkPosts(&data)
// 	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
// }

// func AddParkPosts(c *gin.Context) {

// 	user_id := c.PostForm("user_id")
// 	weixin := c.PostForm("weixin")
// 	posts_type, _ := strconv.Atoi(c.PostForm("posts_type"))
// 	img_url := c.PostForm("img_url")
// 	telephone := c.PostForm("telephone")

// 	data := model.TywParkPostsModel{
// 		UserId:    user_id,
// 		WeiXin:    weixin,
// 		State:     0,
// 		PostsType: posts_type,
// 		ImgUrl:    img_url,
// 		Telephone: telephone,
// 	}
// 	log.Printf("data = %v", data)
// 	log.Printf("c = %v", c)
// 	model.TywCreateParkPosts(&data)
// 	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
// }

// 查询帖子列表
func GetParkPosts(c *gin.Context) {
	size, _ := strconv.Atoi(c.PostForm("size"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	posts_type, _ := strconv.Atoi(c.PostForm("posts_type"))
	user_id := c.PostForm("user_id")
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	if page == 0 {
		page = 1
	}

	data, code := model.TywGetParkPostss(size, page, posts_type, user_id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 编辑帖子
func EditPostsState(c *gin.Context) {
	// username := c.PostForm("username")
	ID, _ := strconv.Atoi(c.PostForm("id"))
	state, _ := strconv.Atoi(c.PostForm("state"))

	// 使用map获取请求参数 接受参数方法与传参方式有很大关系
	// var cate = model.TywParkPostsModel{}
	// c.ShouldBindJSON(&cate)
	code := model.TywEditParkPostsState(ID, state)
	if code == errmsg.SUCCSE {
		response.Success(c, errmsg.GetErrMsg(code), nil)
	} else {
		fmt.Println("编辑帖子失败，修改数据库状态失败")
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
	}
}

// 删除帖子
func DeletePosts(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	id, _ := strconv.Atoi(c.PostForm("id"))
	code := model.DeleteParkPosts(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
