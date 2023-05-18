package v1

import (
	"fmt"
	"log"

	"gindiary/model"
	"gindiary/response"
	"gindiary/util/errmsg"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
注册
*/
func Register(c *gin.Context) {
	//使用map获取请求参数
	// var data model.User
	// _ = c.ShouldBindJSON(&data)
	//获取参数
	username := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// role := c.PostForm("role")
	role, _ := strconv.Atoi(c.PostForm("role"))
	if len(telephone) != 11 {
		response.Fail(c, "手机号必须为11位", nil)
		return
	}
	if len(password) < 6 {
		response.Fail(c, "密码至少为6位", nil)
		return
	}
	if len(username) <= 0 {
		response.Fail(c, "请输入昵称", nil)
		return
	}
	code := model.CheckName(username)
	if code == errmsg.SUCCSE {
		newUser := model.User{
			Username:  username,
			Telephone: telephone,
			Password:  password,
			Role:      role,
		}
		model.CreateUser(&newUser)
		response.Success(c, "注册成功", nil)
	} else {
		response.Fail(c, errmsg.GetErrMsg(code), nil)
	}
}

/*
登录
*/
func Login(c *gin.Context) {

	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	fmt.Printf("postform-telephone:%s\n", telephone)
	fmt.Printf("postform-password:%s\n", password)

	// 使用map获取请求参数
	// var user = model.User{}
	// c.ShouldBindJSON(&user)
	// fmt.Printf("ShouldBindJSON:%s\n", user.Telephone)

	if len(telephone) != 11 {
		response.Fail(c, "手机号必须为11位", nil)

		return
	}
	if len(password) < 4 {
		response.Fail(c, "密码至少为4位", nil)
		return
	}
	code, reUser := model.CheckUser(telephone, password)

	token, err := model.ReleaseToken(reUser)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, "系统异常", nil)
		log.Printf("token generate error: %v", err)
		return
	}
	if code == errmsg.SUCCSE {
		response.Success(c, "登录成功", gin.H{"token": token})
	} else {
		response.Fail(c, errmsg.GetErrMsg(code), nil)
	}
}

/*
获取用户信息
*/
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
		"msg":  errmsg.GetErrMsg(200),
	})
}

/*
退出登录
*/
func Logout(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
		"msg":  errmsg.GetErrMsg(200),
	})
}

/*
编辑用户
*/
func EditUser(c *gin.Context) {
	// username := c.PostForm("username")
	// userId, _ := strconv.Atoi(c.PostForm("id"))

	// 使用map获取请求参数 接受参数方法与传参方式有很大关系
	var user = model.User{}
	c.ShouldBindJSON(&user)

	code := model.CheckName(user.Username)
	if code == errmsg.SUCCSE {
		code2 := model.EditUser(&user)
		if code2 == errmsg.SUCCSE {
			response.Success(c, "修改成功", nil)
		} else {
			response.Fail(c, "保存失败", nil)
		}
	} else {
		response.Fail(c, errmsg.GetErrMsg(code), nil)
	}
}

/*
删除用户
*/
func DeleteUser(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	userId, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteUser(userId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
