package xcxapi

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

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
func TywUserLogin(c *gin.Context) {

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

/*
删除用户
*/
func TywDeleteUser(c *gin.Context) {
	// id, _ := strconv.Atoi(c.PostForm("id"))
	userId, _ := strconv.Atoi(c.PostForm("id"))

	code := model.DeleteUser(userId)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 用户审核
func TywGetSignIn(c *gin.Context) {
	state, _ := strconv.Atoi(c.PostForm("state"))
	// default_community, _ := strconv.Atoi(c.PostForm("default_community"))
	// default_room, _ := strconv.Atoi(c.PostForm("default_room"))
	default_community := c.PostForm("default_community")
	default_room := c.PostForm("default_room")

	fmt.Printf("\n state :%d\n", state)
	fmt.Printf("\n default_community :%s\n", default_community)
	fmt.Printf("\n default_room :%s\n", default_room)

	// 获取上下文中小程序用户信息
	xcxUser := c.Value("user").(model.TywUser)
	xcxUser.State = state
	xcxUser.DefaultCommunity = default_community
	xcxUser.DefaultRoom = default_room

	model.TywEditXcxUserSignIn(&xcxUser)

	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": xcxUser,
		"msg":  "签到成功",
	})
}
