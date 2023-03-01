package model

import (
	"gindiary/util/errmsg"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Username  string `gorm:"type:varchar(20);not null " json:"username"`
	Password  string `gorm:"type:varchar(200);not null " json:"password"`
	Telephone string `gorm:"type:varchar(110);not null" json:"telephone"`
	Role      int    `gorm:"type:int " json:"role"`
}

// 查询用户名是否存在
func CheckName(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERR_USER_USED //用户已存在 1001
	}
	return errmsg.SUCCSE
}

// 注册用户
func CreateUser(data *User) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 编辑用户
func EditUser(data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", data.ID).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}

// 密码加密
func (u *User) BeforeSave() {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err == nil {
		u.Password = string(hasedPassword)
	} else {
		u.Password = "555555"
	}
}

// 校验用户是否存在，密码是否正确
func CheckUser(fromUser *User) (code int) {
	var user User
	db.Where("telephone = ?", fromUser.Telephone).First(&user)

	if user.ID <= 0 {
		return errmsg.ERR_USER_NOT_EXIST //用户不存在
	} else {
		// // 判断加密密码是否正确
		// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(fromUser.Password)); err != nil {
		// 	// response.Response(ctx, http.StatusBadRequest, 400, "密码错误", nil)
		// 	return errmsg.ERR_PASSWORD_WRONG //500
		// } else {
		// 	return errmsg.SUCCSE
		// }
		// 判断密码是否正确
		if user.Password == fromUser.Password {
			return errmsg.SUCCSE
		} else {
			return errmsg.ERR_PASSWORD_WRONG
		}
	}
}
