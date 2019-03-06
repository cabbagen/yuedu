package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"yuedu/schema"
	"yuedu/utils"
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
	CreatedAt         string     `json:"createdAt"`
	Articles          int        `json:"articles"`      // 作品数
	Followers         int        `json:"followers"`     // 粉丝数
	Followings        int        `json:"followings"`    // 关注人员列表
}

func (um UserModel) GetUserInfo(userId int) UserInfo {
	var userInfo UserInfo

	um.database.Table("yd_users").Where("id = ?", userId).Scan(&userInfo)

	um.database.Table("yd_articles").Where("anchor = ?", userId).Count(&userInfo.Articles)

	relationModel := NewRelationModel()

	userInfo.Followers, userInfo.Followings = relationModel.GetUserFollowerCount(userId), relationModel.GetUserFollowingCount(userId)

	return userInfo
}

func (um UserModel) GetUserInfoByName(username string) schema.User {
	var userInfo schema.User

	um.database.Where("username = ?", username).First(&userInfo)

	return userInfo
}

// 新建用户
func (um UserModel) CreateUserInfo(user schema.User) bool {
	var userInfo schema.User

	um.database.Table("yd_users").Where("username = ?", user.Username).First(&userInfo)

	if userInfo.ID > 0 {
		return false
	}

	um.database.Create(&user)

	return true
}

func (um UserModel) ValidateUserInfo(username, password string) bool {
	var userInfo schema.User

	um.database.Table("yd_users").Where("username = ? and password = ?", username, utils.MakeMD5(password)).First(&userInfo)

	return userInfo.ID > 0
}