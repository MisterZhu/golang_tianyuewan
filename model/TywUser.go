package model

import (
	"fmt"
	"gindiary/util/errmsg"
	"time"
)

type TywUser struct {
	ID               uint      `gorm:"primary_key;auto_increment" json:"id"`
	Username         string    `gorm:"type:varchar(20);not null " json:"username"`
	Avatar           string    `gorm:"type:varchar(200);not null " json:"avater"`
	Password         string    `gorm:"type:varchar(200);not null " json:"password"`
	Telephone        string    `gorm:"type:varchar(110);not null" json:"telephone"`
	Role             int       `gorm:"type:int " json:"role"`
	UserId           string    `gorm:"type:varchar(200);not null " json:"user_id"`
	OpenId           string    `gorm:"type:varchar(200);not null " json:"open_id"`
	State            int       `gorm:"type:int " json:"state"`
	DefaultCommunity string    `gorm:"type:varchar(200);not null " json:"default_community"`
	DefaultRoom      string    `gorm:"type:varchar(200);not null " json:"default_room"`
	CreatedAt        time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// 查询用户OpenId是否存在
func TywCheckOpenid(openId string) (code int, reUser TywUser) {

	var user TywUser
	db.Where("open_id = ?", openId).First(&user)
	if user.ID <= 0 {
		return errmsg.ERR_USER_NOT_EXIST, user //用户不存在
	} else {
		return errmsg.SUCCSE, user
	}
}

// 注册用户
func TywCreateXcxUser(data *TywUser) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 编辑用户
func TywEditXcxUserInfo(data *TywUser) int {
	var user TywUser
	var maps = make(map[string]interface{})
	maps["State"] = data.State
	maps["DefaultCommunity"] = data.DefaultCommunity
	maps["DefaultRoom"] = data.DefaultRoom
	maps["Telephone"] = data.Telephone
	fmt.Println("编辑用户")

	err := db.Model(&user).Where("user_id = ?", data.UserId).Updates(maps).Error
	if err != nil {
		fmt.Println("编辑用户失败")

		return errmsg.ERROR
	}
	fmt.Println("编辑用户成功")

	return errmsg.SUCCSE
}

// 编辑用户昵称
func TywEditXcxUserName(data *TywUser) int {
	var user TywUser
	var maps = make(map[string]interface{})
	maps["Username"] = data.Username

	err := db.Model(&user).Where("user_id = ?", data.UserId).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func TywDeleteXcxUser(id string) int {
	var user TywUser
	err = db.Where("user_id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}

// 校验用户是否存在，密码是否正确
func TywCheckXcxUser(telephone string, password string) (code int, reUser TywUser) {
	var user TywUser
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
