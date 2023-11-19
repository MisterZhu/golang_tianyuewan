package v1

import (
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
添加文章
*/
func AddArticle(c *gin.Context) {
	// title := c.PostForm("title")
	// content := c.PostForm("content")
	// img := c.PostForm("img")
	// cid, _ := strconv.Atoi(c.PostForm("cid"))
	// data := model.Article{
	// 	Cid:     cid,
	// 	Img:     img,
	// 	Content: content,
	// 	Title:   title,
	// }

	var data model.Article
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}
	// _ = c.ShouldBindJSON(&data)
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
	var formData FormDataList
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
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

	data, code := model.GetCateArt(formData.ID, formData.Size, formData.Page)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 查询单个文章
func GetArticleInfo(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormDataList
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	data, code := model.GetArtInfo(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArts(c *gin.Context) {
	var formData FormDataList
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}

	data, code := model.GetArts(formData.Size, formData.Page)
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
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormDataList
	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteArticle(formData.ID)

	response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), nil)

}
