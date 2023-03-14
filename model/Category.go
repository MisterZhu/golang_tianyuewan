package model

import (
	"gindiary/util/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(100);not null " json:"name"`
	Content   string `gorm:"type:text;not null" json:"content"`
	ImageUrls string `gorm:"type:text;not null" json:"image_urls"`
}

// 查询分类详情信息
func CheckCategoryDet(id int) (date *Category, code int) {
	/*
		var user User
		db.Where("id = ?", userId).First(&user)
	*/
	var cate Category
	db.Where("id = ?", id).First(&cate)
	if cate.ID > 0 {
		return &cate, errmsg.SUCCSE //分类已存在
	}
	return &cate, errmsg.ERR_CATE_NONE
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERR_CATE_USED //分类已存在
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreateCategory(data *Category) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询分类列表
func GetCates(size int, page int) []Category {
	var cate []Category
	err = db.Limit(size).Offset((page - 1) * size).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

// todo 查询分类下的所有文章

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}
