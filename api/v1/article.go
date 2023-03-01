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
添加文章
*/
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CreateArt(&data)
	// response.Success(c, errmsg.GetErrMsg(code), nil)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章
func GetCateArts(c *gin.Context) {

}

// 查询单个文章
func GetArticle(c *gin.Context) {

}

// 查询分类
func GetArts(c *gin.Context) {
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

	data := model.GetArts(size, page)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 编辑
func EditArticle(c *gin.Context) {
	// username := c.PostForm("username")
	// userId, _ := strconv.Atoi(c.PostForm("id"))

	// 使用map获取请求参数 接受参数方法与传参方式有很大关系
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	id := int(data.ID)

	code := model.EditArt(id, &data)

	response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), nil)

}

// 删除
func DeleteArticle(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	id, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteArticle(id)

	response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), nil)

}
