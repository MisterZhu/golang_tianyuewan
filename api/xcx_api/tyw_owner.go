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
添加业主申请
*/
func AddOwnerApply(c *gin.Context) {

	// community := c.PostForm("community")
	// room := c.PostForm("room")
	// img_url := c.PostForm("img_url")
	// telephone := c.PostForm("telephone")
	// user_id := c.PostForm("user_id")
	var data model.TywOwnerModel
	if err := c.ShouldBindJSON(&data); err != nil {
		// 处理参数绑定错误
		log.Printf("Error binding request data: %v", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}
	data.State = 1

	code := model.TywCreateOwner(&data)
	if code == errmsg.SUCCSE {
		user := model.TywUser{
			State:            1,
			DefaultCommunity: data.Community,
			DefaultRoom:      data.Room,
			UserId:           data.UserId,
		}
		code1 := model.TywEditXcxUserInfo(&user)
		if code1 == errmsg.SUCCSE {
			response.Success(c, errmsg.GetErrMsg(code), nil)
		} else {
			fmt.Println("审核失败，修改数据库个人状态失败")
			response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		}
	} else {
		fmt.Println("审核失败，修改数据库申请信息失败")
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
	}
}

// 查询业主申请
func GetOwnerApply(c *gin.Context) {
	// size, _ := strconv.Atoi(c.PostForm("size"))
	// page, _ := strconv.Atoi(c.PostForm("page"))
	// user_id := c.PostForm("user_id")

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


	data, code := model.TywGetOwners(formData.Size, formData.Page, formData.UserId)
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
	// ID, _ := strconv.Atoi(c.PostForm("id"))
	// state, _ := strconv.Atoi(c.PostForm("state"))
	var formData FormIdData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}

	log.Printf("----------------ID = %v", formData.ID)
	log.Printf("----------------state = %v", formData.State)

	info, code := model.GetOwnersInfo(formData.ID)
	if code == errmsg.SUCCSE {
		info.State = formData.State
		// default_community := c.PostForm("default_community")
		// default_room := c.PostForm("default_room")
		// user_id := c.PostForm("user_id")
		// telephone := c.PostForm("telephone")

		// 使用map获取请求参数 接受参数方法与传参方式有很大关系
		// var cate = model.TywOwnerModel{}
		// c.ShouldBindJSON(&cate)
		log.Printf("----------------info.State = %v", info.State)

		code := model.TywEditOwnerState(formData.ID, formData.State)
		if code == errmsg.SUCCSE {
			user := model.TywUser{
				State:              formData.State,
				DefaultCommunity:   info.Community,
				DefaultRoom:        info.Room,
				UserId:             info.UserId,
				Telephone:          info.Telephone,
				DefaultCommunityId: info.CommunityId,
			}
			code1 := model.TywEditXcxUserInfo(&user)
			if code1 == errmsg.SUCCSE {
				response.Success(c, errmsg.GetErrMsg(code), nil)
			} else {
				fmt.Println("审核失败，修改数据库个人状态失败")
				response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
			}
		} else {
			fmt.Println("审核失败，修改数据库申请信息失败")
			response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		}
	} else {
		fmt.Println("审核失败，修改数据库申请信息失败")
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
	}

}

// 删除申请数据
func DeleteOwner(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteOwner(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
