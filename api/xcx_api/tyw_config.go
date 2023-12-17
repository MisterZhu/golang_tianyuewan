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

type ConfigData struct {
	ID     int    `json:"id"`
	State  string `json:"state"`
	UserId string `json:"user_id"`
}

/*
新增配置字典
*/
func AddConfig(c *gin.Context) {
	// name := c.PostForm("name")
	// detail_name := c.PostForm("detail_name")
	// address := c.PostForm("address")

	var posts model.TywConfigModel
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
	model.TywCreateConfig(&posts)

	response.Success(c, errmsg.GetErrMsg(errmsg.SUCCSE), nil)
}

// 查询配置字典列表
func GetConfig(c *gin.Context) {

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

	data, code := model.TywGetConfigs(formData.Size, formData.Page)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

// 编辑配置字典State
func EdiConfigState(c *gin.Context) {
	// username := c.PostForm("username")
	// ID, _ := strconv.Atoi(c.PostForm("id"))
	// state, _ := strconv.Atoi(c.PostForm("state"))

	var formData ConfigData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.TywEditConfigState(formData.ID, formData.State)
	if code == errmsg.SUCCSE {
		response.Success(c, errmsg.GetErrMsg(code), nil)
	} else {
		fmt.Println("编辑配置字典失败，修改数据库状态失败")
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
	}
}

// 删除配置字典
func DeleteConfig(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteConfig(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

// 获取配置字典详情
func GetDetailConfig(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormConfigData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}

	data, code := model.TywGetConfigInfo(formData.Name)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})

}
