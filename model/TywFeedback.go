package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type TywFeedbackModel struct {
	gorm.Model
	Content string `gorm:"type:varchar(1024);not null " json:"content"`
	UserId  string `gorm:"type:text;not null" json:"user_id"`
}

// 新增申请
func TywCreateFeedback(data *TywFeedbackModel) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询所有申请列表
func TywGetFeedbacks(size, page int) ([]TywFeedbackModel, int) {
	var posts []TywFeedbackModel
	dbQuery := db.Order("created_at desc").Limit(size).Offset((page - 1) * size)

	err := dbQuery.Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return posts, errmsg.SUCCSE
}

// 删除 申请
func DeleteFeedback(id int) int {
	var cate TywFeedbackModel
	err = db.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}
