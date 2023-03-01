package model

import (
	"fmt"
	"gindiary/util/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	ID       uint     `gorm:"primary_key;auto_increment" json:"id"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE

}

// todo 查询分类下的所有文章
// todo 查询单个文章

// 查询文章列表
func GetArts(size int, page int) []Article {
	var art []Article
	err = db.Limit(size).Offset((page - 1) * size).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return art
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	fmt.Printf("%d", id)
	fmt.Printf("title = %s", data.Title)
	fmt.Printf("Content = %s", data.Content)

	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
