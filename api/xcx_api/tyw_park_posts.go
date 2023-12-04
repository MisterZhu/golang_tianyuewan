package xcxapi

import (
	"fmt"
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"
	"log"

	"net/http"

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

// 查询帖子列表
func GetParkPosts(c *gin.Context) {
	var formData FormDataList
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	switch {
	case formData.Size > 100:
		formData.Size = 100
	case formData.Size <= 0:
		formData.Size = 10
	}
	if formData.Page == 0 {
		formData.Page = 1
	}

	data, code := model.TywGetParkPostss(formData.Size, formData.Page, formData.PostsType, formData.UserId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 查询单个帖子
func GetParkInfo(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	data, code := model.TywGetParkPostsInfo(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 编辑帖子
func EditPostsState(c *gin.Context) {
	var formData FormIdData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.TywEditParkPostsState(formData.ID, formData.State)
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
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteParkPosts(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
