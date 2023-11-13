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
添加业主申请
*/
func AddOwnerApply(c *gin.Context) {

	community := c.PostForm("community")
	room := c.PostForm("room")
	//state := c.PostFormArray("state")
	// state, _ := strconv.Atoi(c.PostForm("state"))
	img_url := c.PostForm("img_url")
	telephone := c.PostForm("telephone")
	user_id := c.PostForm("user_id")

	data := model.TywOwnerModel{
		Community: community,
		Room:      room,
		State:     0,
		ImgUrl:    img_url,
		Telephone: telephone,
		UserId:    user_id,
	}
	log.Printf("data = %v", data)
	log.Printf("c = %v", c)
	code := model.TywCreateOwner(&data)
	if code == errmsg.SUCCSE {
		user := model.TywUser{
			State:            1,
			DefaultCommunity: community,
			DefaultRoom:      room,
			UserId:           user_id,
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
	size, _ := strconv.Atoi(c.PostForm("size"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	user_id := c.PostForm("user_id")

	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	if page == 0 {
		page = 1
	}

	data, code := model.TywGetOwners(size, page, user_id)
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
	log.Printf("----------------ID = %v", ID)
	log.Printf("----------------state = %v", state)

	info, code := model.GetOwnersInfo(ID)
	if code == errmsg.SUCCSE {
		info.State = state
		// default_community := c.PostForm("default_community")
		// default_room := c.PostForm("default_room")
		// user_id := c.PostForm("user_id")
		// telephone := c.PostForm("telephone")

		// 使用map获取请求参数 接受参数方法与传参方式有很大关系
		// var cate = model.TywOwnerModel{}
		// c.ShouldBindJSON(&cate)
		log.Printf("----------------info.State = %v", info.State)

		code := model.TywEditOwnerState(ID, state)
		if code == errmsg.SUCCSE {
			user := model.TywUser{
				State:            state,
				DefaultCommunity: info.Community,
				DefaultRoom:      info.Room,
				UserId:           info.UserId,
				Telephone:        info.Telephone,
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
	id, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteOwner(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
