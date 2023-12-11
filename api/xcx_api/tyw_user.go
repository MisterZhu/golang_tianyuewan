package xcxapi

import (
	"encoding/json"
	"fmt"
	"gindiary/model"
	"gindiary/response"
	"gindiary/util"
	"gindiary/util/errmsg"
	"log"

	"github.com/segmentio/ksuid"

	"net/http"

	"github.com/gin-gonic/gin"
)

/*
登录并注册
*/
func TywUserLogin(c *gin.Context) {

	// Code := c.PostForm("code")
	// InviterID := c.PostForm("inviter_id")
	var formData FormCodeData
	// 使用 ShouldBindJSON 解析请求数据到结构体
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	fmt.Printf("\nloginData :%s\n", formData.Code)

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code "
	url = fmt.Sprintf(url, util.AppID, util.AppSecret, formData.Code)

	// 发起请求
	res, _ := http.Get(url)
	// 成功后获取openId
	wxRes := model.WXLoginRes{}
	json.NewDecoder(res.Body).Decode(&wxRes)
	wxRes.OpenId = "oRQ4I42mhyGI76AdXvgbJdQlzo_I"
	fmt.Printf("wxRes.OpenId:%s\n", wxRes.OpenId)
	if len(wxRes.OpenId) <= 0 {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}

	code, reUser := model.TywCheckOpenid(wxRes.OpenId)

	// 用户已存在
	if code == errmsg.SUCCSE {
		fmt.Printf("\n用户已存在\n")
		token, err := model.ReleaseTywToken(reUser)
		if err != nil {
			response.Response(c, http.StatusUnprocessableEntity, 500, "系统异常", nil)
			log.Printf("token generate error: %v", err)
			return
		}
		if code == errmsg.SUCCSE {
			// response.Success(c, "登录成功", gin.H{"token": token})
			c.JSON(http.StatusOK, gin.H{
				"code":  200,
				"token": token,
				"data":  reUser,
				"msg":   errmsg.GetErrMsg(200),
			})
		} else {
			response.Fail(c, errmsg.GetErrMsg(code), nil)
		}

	} else {
		fmt.Printf("\n用户不存在，注册\n")

		// 生成短UUID
		id := ksuid.New().String()[:8]
		fmt.Println(id)
		newUser := model.TywUser{
			Username: "微信用户",
			OpenId:   wxRes.OpenId,
			UserId:   id,
		}
		model.TywCreateXcxUser(&newUser)
		token, err := model.ReleaseTywToken(reUser)
		if err != nil {
			response.Response(c, http.StatusUnprocessableEntity, 500, "系统异常", nil)
			log.Printf("token generate error: %v", err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":  200,
				"token": token,
				"data":  newUser,
				"msg":   errmsg.GetErrMsg(200),
			})
		}

	}
}

// 查询帖子列表
func GetTYWUsers(c *gin.Context) {
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

	data, code := model.TywGetUserList(formData.Size, formData.Page)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
	//response.Response(c, http.StatusOK, 200, errmsg.GetErrMsg(code), data)

}

/*
删除用户
*/
func TywDeleteUser(c *gin.Context) {
	// user_id := c.PostForm("user_id")
	// id, _ := strconv.Atoi(c.PostForm("id"))
	var formData FormIdData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	code := model.DeleteTywUser(formData.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 用户审核
func TywCheckUser(c *gin.Context) {

	var posts model.TywUser
	if err := c.ShouldBindJSON(&posts); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}

	fmt.Printf("\n user_id :%s\n", posts.UserId)
	fmt.Printf("\n state :%d\n", posts.State)
	fmt.Printf("\n default_community :%s\n", posts.DefaultCommunity)
	fmt.Printf("\n default_room :%s\n", posts.DefaultRoom)

	// 获取上下文中小程序用户信息
	tywUser := c.Value("user").(model.TywUser)
	tywUser.State = posts.State
	tywUser.DefaultCommunity = posts.DefaultCommunity
	tywUser.DefaultRoom = posts.DefaultRoom
	tywUser.UserId = posts.UserId

	model.TywEditXcxUserInfo(&tywUser)

	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": tywUser,
		"msg":  "审核通过",
	})
}

// 用户修改昵称
func TywChangeUserName(c *gin.Context) {
	var posts model.TywUser
	if err := c.ShouldBindJSON(&posts); err != nil {
		log.Printf("Error binding request data: %v", err)
		c.JSON(400, gin.H{"message": "Invalid request data"})
		return
	}
	fmt.Printf("\n posts.Username :%s\n", posts.Username)
	// 获取上下文中小程序用户信息
	tywUser := c.Value("user").(model.TywUser)
	tywUser.Username = posts.Username

	model.TywEditXcxUserName(&tywUser)

	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": tywUser,
		"msg":  "修改成功",
	})
}
