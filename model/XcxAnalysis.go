package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type XcxAnalyModel struct {
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

// 新增解析
func CreateXcxAnaly(data *XcxAnalyModel) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询解析列表
func GetAnalys(size int, page int, userId string) []XcxAnalyModel {

	var cate []XcxAnalyModel
	// err = db.Limit(size).Offset(page * size).Find(&cate).Error
	err = db.Order("updated_at desc").Where("user_id = ?", userId).Limit(size).Offset(page * size).Find(&cate).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}
