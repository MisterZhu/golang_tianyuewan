package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type TywConfigModel struct {
	gorm.Model
	Title string `gorm:"type:varchar(110);not null" json:"title"`
	State string `gorm:"type:varchar(110);not null" json:"state"`
	Name  string `gorm:"type:text;not null" json:"name"`
}

// 新增配置字典
func TywCreateConfig(data *TywConfigModel) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询配置字典列表
func TywGetConfigs(size, page int) ([]TywConfigModel, int) {
	var posts []TywConfigModel
	dbQuery := db.Order("updated_at asc").Limit(size).Offset(page * size)

	err := dbQuery.Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return posts, errmsg.SUCCSE
}

// todo 查询单个字典
func TywGetConfigInfo(name string) (TywConfigModel, int) {
	var art TywConfigModel
	err := db.Where("name = ?", name).First(&art).Error
	if err != nil {
		return art, errmsg.ERR_ART_NONE
	}
	return art, errmsg.SUCCSE

}

// 编辑配置字典
func TywEditConfigState(id int, newState string) int {
	var art TywConfigModel
	var maps = make(map[string]interface{})
	maps["State"] = newState

	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除 配置字典
func DeleteConfig(id int) int {
	var cate TywConfigModel
	err = db.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}
