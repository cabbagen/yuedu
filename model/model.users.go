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

func (um UserModel) GetUserInfo(userId int) (UserInfo, error) {
	var userInfo UserInfo

	if result := um.database.Table("yd_users").Where("id = ?", userId).Scan(&userInfo); result.Error != nil {
		return userInfo, result.Error
	}

	if result := um.database.Table("yd_articles").Where("anchor = ?", userId).Count(&userInfo.Articles); result.Error != nil {
		return userInfo, result.Error
	}

	if followers, error := NewRelationModel().GetUserFollowerCount(userId); error != nil {
		return userInfo, error
	} else {
		userInfo.Followers = followers
	}

	if followings, error := NewRelationModel().GetUserFollowingCount(userId); error != nil {
		return userInfo, error
	} else {
		userInfo.Followings = followings
	}

	return userInfo, nil
}

func (um UserModel) GetUserInfoByName(username string) (schema.User, error) {
	var userInfo schema.User

	if result := um.database.Where("username = ?", username).First(&userInfo); result.Error != nil {
		return userInfo, result.Error
	}

	return userInfo, nil
}

func (um UserModel) GetUserInfoByUserIds(userIds []int) ([]schema.User, error) {
	var users []schema.User

	if result := um.database.Table("yd_users").Where("id in (?)", userIds).Find(&users); result.Error != nil {
		return users, nil
	}

	return users, nil
}

// 新建用户
func (um UserModel) CreateUserInfo(user schema.User) error {
	if result := um.database.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (um UserModel) ValidateUserInfo(username, password string) (bool, error) {
	var userInfo schema.User

	result := um.database.Table("yd_users").
		Where("username = ? and password = ?", username, utils.MakeMD5(password)).
		First(&userInfo)

	if result.Error != nil {
		return false, result.Error
	}

	return userInfo.ID > 0, nil
}

// 获取用户的额外信息
func (um UserModel) GetUserExtraInfo(userId, articleId int) (map[string]interface{}, error) {
	var extraInfo map[string]interface{} = make(map[string]interface{})

	var article schema.Article

	articleResult := um.database.Table("yd_articles").Where("id = ?", articleId).First(&article)

	if articleResult.Error != nil {
		return extraInfo, articleResult.Error
	}

	// 查询该用户是否收藏了文章
	isCollected, error := NewCollectionModel().IsCollectedArticle(userId, articleId)

	if error != nil {
		return extraInfo, error
	}

	extraInfo["isCollected"] = isCollected

	// 用户是否喜欢该文章
	isSupported, error := NewSupportModel().IsSupportArticle(userId, articleId)

	if error != nil {
		return extraInfo, error
	}

	extraInfo["isSupported"] = isSupported

	// 用户是否为文章主播的粉丝
	followFlag, error := NewRelationModel().GetUsersRelations(userId, article.Anchor)

	if error != nil {
		return extraInfo, error
	}

	extraInfo["isFollowing"] = followFlag == 1

	return extraInfo, nil
}