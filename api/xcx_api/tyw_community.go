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
新增社区
*/
func AddCommunity(c *gin.Context) {
	var posts model.TywCommunityModel
	if err := c.ShouldBindJSON(&posts); err != nil {
		// 处理参数绑定错误
		log.Printf("Error binding request data: %v", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}

	// 可以在这里对 posts 进行一些其他处理，然后插入数据库等操作
	log.Printf("posts = %+v", posts)
	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
}

// 查询社区列表
func GetCommunity(c *gin.Context) {
	size, _ := strconv.Atoi(c.PostForm("size"))
	page, _ := strconv.Atoi(c.PostForm("page"))

	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	if page == 0 {
		page = 1
	}

	data, code := model.TywGetCommunitys(size, page)
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
	ID, _ := strconv.Atoi(c.PostForm("id"))
	state, _ := strconv.Atoi(c.PostForm("state"))

	// 使用map获取请求参数 接受参数方法与传参方式有很大关系
	// var cate = model.TywCommunityModel{}
	// c.ShouldBindJSON(&cate)
	code := model.TywEditCommunityState(ID, state)
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
	id, _ := strconv.Atoi(c.PostForm("id"))
	code := model.DeleteCommunity(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
