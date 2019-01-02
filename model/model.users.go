package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	database        *gorm.DB
}

func NewUserModel() UserModel {
	return UserModel { database.GetDataBase() }
}


// 用户信息实体
type UserInfo struct {
	Id                int        `json:"id"`
	Username          string     `json:"username"`
	Password          string     `json:"password"`
	Gender            int        `json:"gender"`
	Email             string     `json:"email"`
	Address           int        `json:"address"`
	Homepages         string     `json:"homepages"`
	Avatar            string     `json:"avatar"`
	Backdrop          string     `json:"backdrop"`
	Extra             string     `json:"extra"`
	Articles          int        `json:"articles"`    // 作品数
	Flowers           int        `json:"flowers"`     // 粉丝数
}

func (um UserModel) GetUserInfo(userId int) UserInfo {
	var userInfo UserInfo

	um.database.Table("yd_users").Where("id = ?", userId).Scan(&userInfo)

	um.database.Table("yd_articles").Where("anchor = ?", userId).Count(&userInfo.Articles)

	um.database.Table("yd_relations").Where("relation_user_id = ? and relation_type != 1", userId).Count(&userInfo.Flowers)

	return userInfo
}
