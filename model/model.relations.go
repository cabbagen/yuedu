package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"yuedu/schema"
)

type RelationModel struct {
	database        *gorm.DB
}

func NewRelationModel() RelationModel {
	return RelationModel { database.GetDataBase() }
}

// 获取用户的粉丝
// 即 用户被其他人关注 or 用户和粉丝之间相互关注
func (rm RelationModel) getUserFollowers(userId int) []schema.Relation {
	var relations []schema.Relation

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 3 ) or ( user_id = ? and relation_type = 1 )"

	rm.database.Table("yd_relations").
		Select("id, user_id, relation_user_id, relation_type").
		Where(condition, userId, userId, userId).
		Find(&relations)

	return relations
}

// 获取用户的关注的人员列表
// 即 用户关注的人 or 用户和粉丝之间的相互关注
func (rm RelationModel) getUserFollowings(userId int) []schema.Relation {
	var relations []schema.Relation

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 3 ) or ( relation_user_id = ? and relation_type = 2 )"

	// 互相关注
	rm.database.Table("yd_relations").
		Select("id, user_id, relation_user_id, relation_type").
		Where(condition, userId, userId, userId).
		Find(&relations)


	return relations
}

