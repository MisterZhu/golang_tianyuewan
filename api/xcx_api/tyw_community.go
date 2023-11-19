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

type FormDataList struct {
	Size      int    `json:"size"`
	Page      int    `json:"page"`
	UserId    string `json:"user_id"`
	PostsType int    `json:"posts_type"`
}
type FormIdData struct {
	ID     int    `json:"id"`
	State  int    `json:"state"`
	UserId string `json:"user_id"`
}
type FormCodeData struct {
	Code string `json:"code"`
}

/*
新增社区
*/
func AddCommunity(c *gin.Context) {
	// name := c.PostForm("name")
	// detail_name := c.PostForm("detail_name")
	// address := c.PostForm("address")

	var posts model.TywCommunityModel
	// posts.Name = name
	// posts.DetailName = detail_name
	// posts.Address = address
	if err := c.ShouldBindJSON(&posts); err != nil {
		// 处理参数绑定错误
		log.Printf("Error binding request data: %v", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}

	// 可以在这里对 posts 进行一些其他处理，然后插入数据库等操作
	log.Printf("posts = %+v", posts)
	model.TywCreateCommunity(&posts)

	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
}

// 查询社区列表
func GetCommunity(c *gin.Context) {

	//前端 'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
	//解析方法
	// size, _ := strconv.Atoi(c.PostForm("size"))
	// page, _ := strconv.Atoi(c.PostForm("page"))

	//前端 "Content-Type": "application/json"
	//解析方法
	var formData FormDataList
	// 使用 ShouldBindJSON 解析请求数据到结构体
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

	data, code := model.TywGetCommunitys(formData.Size, formData.Page)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 编辑社区State
func EdiCommunityState(c *gin.Context) {
	// username := c.PostForm("username")
	// ID, _ := strconv.Atoi(c.PostForm("id"))
	// state, _ := strconv.Atoi(c.PostForm("state"))

	var formData FormIdData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.TywEditCommunityState(formData.ID, formData.State)
	if code == errmsg.SUCCSE {
		response.Success(c, errmsg.GetErrMsg(code), nil)
	} else {
		fmt.Println("编辑社区失败，修改数据库状态失败")
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
	}
}

// 删除社区
func DeleteCommunity(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteCommunity(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
