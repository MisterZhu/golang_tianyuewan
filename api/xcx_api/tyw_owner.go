package xcxapi

import (
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"
	"log"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
添加业主申请
*/
func AddOwnerApply(c *gin.Context) {

	community := c.PostForm("community")
	room := c.PostForm("room")
	//state := c.PostFormArray("state")
	// state, _ := strconv.Atoi(c.PostForm("state"))
	img_url := c.PostForm("img_url")

	telephone := c.PostForm("telephone")

	data := model.TywOwnerModel{
		Community: community,
		Room:      room,
		State:     0,
		ImgUrl:    img_url,
		Telephone: telephone,
	}
	log.Printf("data = %v", data)
	log.Printf("c = %v", c)
	model.TywCreateOwner(&data)
	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)

}

// 查询业主申请
func GetOwnerApply(c *gin.Context) {
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

	data := model.TywGetOwners(size, page)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// // 查询分类详情
// func GetCategoryDet(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.PostForm("id"))
// 	// id := c.PostForm("id")

// 	data, code := model.CheckCategoryDet(id)
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": code,
// 		"data": data,
// 		"msg":  errmsg.GetErrMsg(code),
// 	})
// 	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

// }

// 编辑申请状态
func EditApplyState(c *gin.Context) {
	// username := c.PostForm("username")
	ID, _ := strconv.Atoi(c.PostForm("id"))
	state, _ := strconv.Atoi(c.PostForm("state"))

	// 使用map获取请求参数 接受参数方法与传参方式有很大关系
	// var cate = model.TywOwnerModel{}
	// c.ShouldBindJSON(&cate)

	code := model.TywEditOwnerState(ID, state)
	response.Success(c, errmsg.GetErrMsg(code), nil)

}

// 删除申请数据
func DeleteOwner(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	userId, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteOwner(userId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
