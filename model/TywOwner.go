package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type TywOwnerModel struct {
	gorm.Model
	OpenId        string `gorm:"type:varchar(1024);not null " json:"open_id"`
	UserId        string `gorm:"type:varchar(200);not null " json:"user_id"`
	OriginURL     string `gorm:"type:text;not null" json:"origin_url"`
	Title         string `gorm:"type:text;not null" json:"title"`
	Cover         string `gorm:"type:text;not null" json:"cover"`
	DownloadImage string `gorm:"type:text;not null" json:"download_image"`
	Video         string `gorm:"type:text;not null" json:"video"`
	URL           string `gorm:"type:text;not null" json:"url"`
	Down          string `gorm:"type:text;not null" json:"down"`
	Images        string `gorm:"type:text;not null" json:"images"`
	BigFile       bool   `gorm:"default:false" json:"big_file"`
}

// 新增申请
func TywCreateOwner(data *TywOwnerModel) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询申请列表
func TywGetOwners(size int, page int, userId string) []TywOwnerModel {

	var cate []TywOwnerModel
	// err = db.Limit(size).Offset((page - 1) * size).Find(&cate).Error
	err = db.Order("updated_at desc").Where("user_id = ?", userId).Limit(size).Offset((page - 1) * size).Find(&cate).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类
func TywEditOwner(id int, data *TywOwnerModel) int {
	var cate TywOwnerModel
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}
