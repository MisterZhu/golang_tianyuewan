package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type TywOwnerModel struct {
	gorm.Model
	CommunityId int    `gorm:"type:int " json:"community_id"`
	Community   string `gorm:"type:varchar(1024);not null " json:"community"`
	Room        string `gorm:"type:varchar(200);not null " json:"room"`
	State       int    `gorm:"type:int " json:"state"`
	UserId      string `gorm:"type:text;not null" json:"user_id"`
	Telephone   string `gorm:"type:varchar(110);not null" json:"telephone"`
	ImgUrl      string `gorm:"type:text;not null" json:"img_url"`
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

// 查询所有申请列表
func TywGetOwners(size, page int, userID string) ([]TywOwnerModel, int) {
	var posts []TywOwnerModel
	dbQuery := db.Order("updated_at asc").Limit(size).Offset(page * size)

	// 根据 userID 进行过滤
	if userID != "" {
		dbQuery = dbQuery.Where("user_id = ?", userID)
	}

	err := dbQuery.Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return posts, errmsg.SUCCSE
}

// todo 查询申请详情
func GetOwnersInfo(id int) (TywOwnerModel, int) {
	var art TywOwnerModel
	err := db.Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERR_ART_NONE
	}
	return art, errmsg.SUCCSE

}

// 查询所有申请列表
// func TywGetOwners(size int, page int) []TywOwnerModel {

// 	var cate []TywOwnerModel
// 	// err = db.Limit(size).Offset(page * size).Find(&cate).Error
// 	err = db.Order("updated_at desc").Limit(size).Offset(page * size).Find(&cate).Error

// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil
// 	}
// 	return cate
// }

// 查询分类列表
//
//	func GetCates(size int, page int) []Category {
//		var cate []Category
//		err = db.Limit(size).Offset(page * size).Find(&cate).Error
//		if err != nil && err != gorm.ErrRecordNotFound {
//			return nil
//		}
//		return cate
//	}
//
// 审核申请
// func TywEditOwner(id int, data *TywOwnerModel) int {
// 	var cate TywOwnerModel
// 	var maps = make(map[string]interface{})
// 	maps["State"] = data.State
// 	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
// 	if err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCSE

// }
// 编辑申请信息
func TywEditOwnerState(id int, newState int) int {
	var art TywOwnerModel
	var maps = make(map[string]interface{})
	maps["State"] = newState

	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除 申请
func DeleteOwner(id int) int {
	var cate TywOwnerModel
	err = db.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}
