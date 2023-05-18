package xcxapi

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"gindiary/model"
	"gindiary/response"
	"gindiary/util"
	"gindiary/util/errmsg"

	"github.com/segmentio/ksuid"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
登录并注册
*/
func XcxUserLogin(c *gin.Context) {

	Code := c.PostForm("code")
	InviterID := c.PostForm("inviter_id")

	fmt.Printf("\nloginData :%s\n", Code)

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code "
	url = fmt.Sprintf(url, util.AppID, util.AppSecret, Code)

	// 发起请求
	res, _ := http.Get(url)
	// 成功后获取openId
	wxRes := model.WXLoginRes{}
	json.NewDecoder(res.Body).Decode(&wxRes)
	fmt.Printf("wxRes.OpenId:%s\n", wxRes.OpenId)
	if len(wxRes.OpenId) <= 0 {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERROR), nil)
		return
	}

	code, reUser := model.CheckOpenid(wxRes.OpenId)

	log.Printf("reUserQueryCount: %v\n", reUser.QueryCount)

	// 用户已存在
	if code == errmsg.SUCCSE {
		fmt.Printf("\n用户已存在\n")

		token, err := model.ReleaseXcxToken(reUser)
		if err != nil {
			response.Response(c, http.StatusUnprocessableEntity, 500, "系统异常", nil)
			log.Printf("token generate error: %v", err)
			return
		}
		if code == errmsg.SUCCSE {
			now := time.Now()
			samedate := isSameDay(now, reUser.UpdatedAt)
			if !samedate {
				log.Printf("不是同一天")

				if reUser.QueryCount < 3 {
					reUser.QueryCount = 3
				}
				signSamedate := isSameDay(now, reUser.SiginTime)
				if !signSamedate {
					reUser.SiginCount = 0
					reUser.SiginReward = 2
				}

				model.EditXcxUserQueryCount(&reUser)
			}
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
		newUser := model.XcxUser{
			Username:    "微信用户",
			OpenId:      wxRes.OpenId,
			QueryCount:  3,
			UserId:      id,
			SiginCount:  0,
			SiginReward: 2,
			InviterID:   InviterID,
		}
		model.CreateXcxUser(&newUser)
		token, err := model.ReleaseXcxToken(reUser)
		if err != nil {
			response.Response(c, http.StatusUnprocessableEntity, 500, "系统异常", nil)
			log.Printf("token generate error: %v", err)
			return
		}
		nviterCode, inviterUser := model.CheckOpenid(InviterID)
		if nviterCode == errmsg.SUCCSE {
			InvitedAry := inviterUser.InvitedUsers
			iscontains := strings.Contains(InvitedAry, InviterID)
			if !iscontains {
				if len(InvitedAry) == 0 {
					InvitedAry = InviterID
				} else {
					InvitedAry = InvitedAry + "," + InviterID
				}
				inviterUser.InvitedUsers = InvitedAry
				inviterUser.QueryCount += 10 * (len(InvitedAry) + 1)
				model.EditXcxUserInvited(&inviterUser)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"token": token,
			"data":  newUser,
			"msg":   errmsg.GetErrMsg(200),
		})
	}
}
func isSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
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

// 用户签到
func XcxGetSignIn(c *gin.Context) {
	sigin_count, _ := strconv.Atoi(c.PostForm("sigin_count"))
	sigin_reward, _ := strconv.Atoi(c.PostForm("sigin_reward"))
	fmt.Printf("\n sigin_count :%d\n", sigin_count)
	fmt.Printf("\n sigin_reward :%d\n", sigin_reward)

	switch {
	case sigin_count == 0:
		sigin_reward = 2
		sigin_count = 1

	case sigin_count == 1:
		sigin_reward = 2
		sigin_count = 2

	case sigin_count == 2:
		sigin_reward = 3
		sigin_count = 3

	case sigin_count == 3:
		sigin_reward = 3
		sigin_count = 4

	case sigin_count == 4:
		sigin_reward = 5
		sigin_count = 5

	case sigin_count == 5:
		sigin_reward = 5
		sigin_count = 6

	case sigin_count == 6:
		sigin_reward = 7
		sigin_count = 7

	case sigin_count == 7:
		sigin_reward = 2
		sigin_count = 1

	}
	now := time.Now()
	// 获取上下文中小程序用户信息
	xcxUser := c.Value("user").(model.XcxUser)
	xcxUser.SiginCount = sigin_count
	xcxUser.SiginReward = sigin_reward
	xcxUser.SiginTime = now
	xcxUser.QueryCount = xcxUser.QueryCount + sigin_reward

	model.EditXcxUserSignIn(&xcxUser)

	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": xcxUser,
		"msg":  "签到成功",
	})
}
