package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type TywCommunityModel struct {
	gorm.Model
	Latitude   float64 `gorm:"type:double" json:"latitude"`
	Longitude  float64 `gorm:"type:double" json:"longitude"`
	State      int     `gorm:"type:int " json:"state"`
	UserId     string  `gorm:"type:text;not null" json:"user_id"`
	Address    string  `gorm:"type:varchar(110);not null" json:"address"`
	DetailName string  `gorm:"type:varchar(110);not null" json:"detail_name"`
	Name       string  `gorm:"type:text;not null" json:"name"`
	ImgUrl     string  `gorm:"type:text;not null" json:"img_url"`
}

// 新增社区
func TywCreateCommunity(data *TywCommunityModel) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询社区列表
func TywGetCommunitys(size, page int) ([]TywCommunityModel, int) {
	var posts []TywCommunityModel
	dbQuery := db.Order("created_at desc").Limit(size).Offset((page - 1) * size)

	err := dbQuery.Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return posts, errmsg.SUCCSE
}

// 编辑社区信息
func TywEditCommunityState(id int, newState int) int {
	var art TywCommunityModel
	var maps = make(map[string]interface{})
	maps["State"] = newState

	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除 社区
func DeleteCommunity(id int) int {
	var cate TywCommunityModel
	err = db.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}
