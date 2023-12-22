package model

import (
	"gindiary/util/errmsg"
	// "time"
	"gorm.io/gorm"
)

type TywShopsModel struct {
	gorm.Model
	InMaintenance bool   `gorm:"type:bool" json:"in_maintenance"`
	Negotiable    bool   `gorm:"type:bool" json:"negotiable"`
	State         int    `gorm:"type:int " json:"state"`
	PostsType     int    `gorm:"type:int " json:"posts_type"`
	UserId        string `gorm:"type:text;not null" json:"user_id"`
	Telephone     string `gorm:"type:varchar(110);not null" json:"telephone"`
	WeiXin        string `gorm:"type:varchar(110);not null" json:"wei_xin"`
	Title         string `gorm:"type:text;not null" json:"title"`
	ImgUrl        string `gorm:"type:text;not null" json:"img_url"`
	AnnualRent    string `gorm:"type:text;not null" json:"annual_rent"`
	CommunityId   int    `gorm:"type:int " json:"community_id"`
	Address       string `gorm:"type:text;not null" json:"address"`
}

// 新增帖子
func TywCreateShops(data *TywShopsModel) int {
	// data.BeforeSave()
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

/*
// 仅检索所有类型的帖子
allPosts := TywGetShopss(10, 1, 0, "")
// 检索指定类型和用户的帖子
userPosts := TywGetShopss(10, 1, 1, "specificUserID")
// 仅检索指定用户的帖子
specificUserPosts := TywGetShopss(10, 1, 0, "specificUserID")
*/

// 查询所有帖子列表（postType可传，不传就是查询所有的帖子，postType=1是出租车位帖子，postType=2为求租帖子）
func TywGetShopss(size, page, postType int) ([]TywShopsModel, int) {
	var posts []TywShopsModel
	dbQuery := db.Order("updated_at asc").Limit(size).Offset(page * size)

	// 根据 postType 进行过滤 1：家政 2：装修 3：废品回收
	if postType != 0 {
		dbQuery = dbQuery.Where("posts_type = ?", postType)
	}

	err := dbQuery.Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return posts, errmsg.SUCCSE
}

// todo 查询单个文章
func TywGetShopsInfo(id int) (TywShopsModel, int) {
	var art TywShopsModel
	err := db.Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERR_ART_NONE
	}
	return art, errmsg.SUCCSE

}

// 编辑帖子信息
func TywEditShopsState(id int, newState int) int {
	var art TywShopsModel
	var maps = make(map[string]interface{})
	maps["State"] = newState

	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除 帖子
func DeleteShops(id int) int {
	var cate TywShopsModel
	err = db.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR //
	}
	return errmsg.SUCCSE
}
