package v1

import (
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
添加分类
*/
func AddCategory(c *gin.Context) {

	// name := c.PostForm("title")
	// content := c.PostForm("content")
	// imageUrlsAry := c.PostFormArray("imageUrls")

	var data model.Category
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	// image_urls := strings.Join(formData.ImageUrls, ",")

	// data := model.Category{
	// 	Name:      formData.Name,
	// 	Content:   formData.Content,
	// 	ImageUrls: image_urls,
	// }
	log.Printf("imageUrlsAry = %v", data.ImageUrls)

	log.Printf("data = %v", data)
	log.Printf("c = %v", c)

	// var data model.Category
	// _ = c.ShouldBindJSON(&data)
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

	data := model.GetCates(formData.Size, formData.Page)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 查询分类详情
func GetCategoryDet(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	data, code := model.CheckCategoryDet(formData.ID)
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
	c.ShouldBindJSON(&cate)

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
	//userId, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteCategory(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
