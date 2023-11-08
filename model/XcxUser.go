package model

import (
	"gindiary/util/errmsg"
	"time"
)

// 封装code2session接口返回数据
type WXLoginRes struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type XcxUser struct {
	ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"type:varchar(20);not null " json:"username"`
	Avatar     string    `gorm:"type:varchar(200);not null " json:"avater"`
	Password   string    `gorm:"type:varchar(200);not null " json:"password"`
	Telephone  string    `gorm:"type:varchar(110);not null" json:"telephone"`
	Role       int       `gorm:"type:int " json:"role"`
	UserId     string    `gorm:"type:varchar(200);not null " json:"user_id"`
	QueryCount int       `gorm:"type:int" json:"query_count"`
	OpenId     string    `gorm:"type:varchar(200);not null " json:"open_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`

	SiginCount  int       `gorm:"type:int " json:"sigin_count"`
	SiginReward int       `gorm:"type:int " json:"sigin_reward"`
	SiginTime   time.Time `gorm:"type:datetime;null" json:"sigin_time,omitempty"`

	InviterID    string `gorm:"type:varchar(200);not null" json:"inviter_id"`
	InvitedUsers string `gorm:"type:text;not null" json:"invited_users"`
}

// 查询用户OpenId是否存在
func CheckOpenid(openId string) (code int, reUser XcxUser) {

	var user XcxUser
	db.Where("open_id = ?", openId).First(&user)
	if user.ID <= 0 {
		return errmsg.ERR_USER_NOT_EXIST, user //用户不存在
	} else {
		return errmsg.SUCCSE, user
	}
}

// 注册用户
func CreateXcxUser(data *XcxUser) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 更新用户每日次数
func EditXcxUserQueryCount(data *XcxUser) int {
	var user XcxUser
	var maps = make(map[string]interface{})
	maps["query_count"] = data.QueryCount
	err := db.Model(&user).Where("id = ?", data.ID).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 更新用户邀请数据
func EditXcxUserInvited(data *XcxUser) int {
	var user XcxUser
	var maps = make(map[string]interface{})
	maps["QueryCount"] = data.QueryCount
	maps["InvitedUsers"] = data.InvitedUsers

	err := db.Model(&user).Where("id = ?", data.ID).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 编辑用户
func EditXcxUserSignIn(data *XcxUser) int {
	var user XcxUser
	var maps = make(map[string]interface{})
	maps["SiginCount"] = data.SiginCount
	maps["SiginReward"] = data.SiginReward
	maps["SiginTime"] = data.SiginTime

	err := db.Model(&user).Where("id = ?", data.ID).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteXcxUser(id int) int {
	var user XcxUser
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}

// 校验用户是否存在，密码是否正确
func CheckXcxUser(telephone string, password string) (code int, reUser XcxUser) {
	var user XcxUser
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID <= 0 {
		return errmsg.ERR_USER_NOT_EXIST, user //用户不存在
	} else {
		// 判断密码是否正确
		if user.Password == password {
			return errmsg.SUCCSE, user
		} else {
			return errmsg.ERR_PASSWORD_WRONG, user
		}
	}
}
