package v1

import (
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
添加分类
*/
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateCategory(&data)
		response.Success(c, errmsg.GetErrMsg(code), nil)
	} else {
		response.Fail(c, errmsg.GetErrMsg(code), nil)

	}

}

// 查询分类
func GetCategory(c *gin.Context) {
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

	data := model.GetCates(size, page)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 编辑分类名称
func EditCategory(c *gin.Context) {
	// username := c.PostForm("username")
	// userId, _ := strconv.Atoi(c.PostForm("id"))

	// 使用map获取请求参数 接受参数方法与传参方式有很大关系
	var cate = model.Category{}
	c.Bind(&cate)

	code := model.CheckCategory(cate.Name)
	if code == errmsg.SUCCSE {
		code2 := model.EditCategory(int(cate.ID), &cate)
		response.Success(c, errmsg.GetErrMsg(code2), nil)

	} else {
		response.Fail(c, errmsg.GetErrMsg(code), nil)
	}
}

// 删除
func DeleteCategory(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	userId, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteCategory(userId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
