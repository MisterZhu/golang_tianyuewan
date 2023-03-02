package v1

import (
	"fmt"

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
	username := c.PostForm("username")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// role := c.PostForm("role")
	role, _ := strconv.Atoi(c.PostForm("role"))
	if len(telephone) != 11 {
		response.Fail(c, "手机号必须为11位", nil)
		return
	}
	if len(password) < 4 {
		response.Fail(c, "密码至少为4位", nil)
		return
	}
	if len(username) <= 0 {
		response.Fail(c, "请输入昵称", nil)
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

	ph := c.PostForm("telephone")
	fmt.Printf("postform:%s\n", ph)

	// 使用map获取请求参数
	var user = model.User{}
	c.ShouldBindJSON(&user)
	fmt.Printf("ShouldBindJSON:%s\n", user.Telephone)

	if len(user.Telephone) != 11 {
		response.Fail(c, "手机号必须为11位", nil)

		return
	}
	if len(user.Password) < 4 {
		response.Fail(c, "密码至少为4位", nil)
		return
	}
	code := model.CheckUser(&user)
	if code == errmsg.SUCCSE {
		//用户存在
		response.Success(c, "登录成功", nil)
	} else {
		response.Fail(c, errmsg.GetErrMsg(code), nil)
	}
}

// 编辑用户
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

// 删除用户
func DeleteUser(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	userId, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteUser(userId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}
