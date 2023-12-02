package xcxapi

import (
	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
新增反馈
*/
func AddFeedback(c *gin.Context) {

	var posts model.TywFeedbackModel

	if err := c.ShouldBindJSON(&posts); err != nil {
		// 处理参数绑定错误
		log.Printf("Error binding request data: %v", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}

	// 可以在这里对 posts 进行一些其他处理，然后插入数据库等操作
	log.Printf("posts = %+v", posts)
	model.TywCreateFeedback(&posts)

	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
}

// 查询反馈列表
func GetFeedback(c *gin.Context) {

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
	data, code := model.TywGetFeedbacks(formData.Size, formData.Page)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 删除反馈
func DeleteFeedback(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteFeedback(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
