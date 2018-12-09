package model

import (
	"yuedu/database"
	"yuedu/schema"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	database        *gorm.DB
}

func NewUserModel() UserModel {
	return UserModel { database.GetDataBase() }
}

type FullUserInfo struct {
	schema.User
	Articles             int       `json:"articles"`
	Flowers              int       `json:"flowers"`
}
func (um UserModel) GetFullUserInfo(userId int, fullUserInfo *FullUserInfo) {
	um.database.Table("yd_users").Where("yd_users.id = ?", userId).Scan(fullUserInfo)
	um.database.Table("yd_users").
		Select("count(yd_articles.id) as articles, count(yd_relations.id) as flowsers").
		Where("yd_users.id = ?", userId).
		Joins("left join yd_articles on yd_articles.anchor = yd_users.id").
		Joins("left join yd_relations on yd_relations.relation_user_id = yd_users.id and yd_relations.relation_type != 1").
		Group("yd_articles.anchor, yd_relations.relation_user_id").
		Scan(fullUserInfo)
}
