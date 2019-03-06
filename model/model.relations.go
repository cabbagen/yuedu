package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"yuedu/schema"
	"time"
)

type RelationModel struct {
	database        *gorm.DB
}

func NewRelationModel() RelationModel {
	return RelationModel { database.GetDataBase() }
}

// 查询用户的关注粉丝列表
func (rm RelationModel) GetUserFollowers(userId int) []schema.User {
	var followers []schema.User

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( relation_user_id = ? and relation_type = 1 )"

	rm.database.Table("yd_relations").
		Select("yd_users.*").
		Where(condition, userId, userId, userId).
		Joins("inner join yd_users on yd_users.id = yd_relations.relation_user_id").
		Find(&followers)

	return followers
}

// 查询粉丝的数量
func (rm RelationModel) GetUserFollowerCount(userId int) int {
	var count int
	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( relation_user_id = ? and relation_type = 1 )"

	rm.database.Table("yd_relations").Where(condition, userId, userId, userId).Count(&count)

	return count
}

// 查询用户正在关注的人员的列表
func (rm RelationModel) GetUserFollowings(userId int) []schema.User {
	var followings []schema.User

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( user_id = ? and relation_type = 1 )"

	// 互相关注
	rm.database.Table("yd_relations").
		Select("yd_users.*").
		Where(condition, userId, userId, userId).
		Joins("inner join yd_users on yd_users.id = yd_relations.relation_user_id").
		Find(&followings)

	return followings
}

func (rm RelationModel) GetUserFollowingCount(userId int) int {
	var count int

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( user_id = ? and relation_type = 1 )"

	rm.database.Table("yd_relations").Where(condition, userId, userId, userId).Count(&count)

	return count
}

// 添加关注
func (rm RelationModel) CreateUserFollowing(userId, relationUserId int) {
	var relation schema.Relation

	var condition string = "user_id = ? and relation_user_id = ? and relation_type = 1"

	// 查询该用户是否被关注过
	// 如果被关注过，则修改这条记录
	// 如果未被关注过，则新建一条记录
	rm.database.Table("yd_relations").Where(condition, relationUserId, userId).First(&relation)

	if relation.ID > 0 {
		rm.UpdateUserRelation(relation.ID, map[string]interface{} {"relation_type": 2})
		return
	}

	rm.CreateUserRelation(userId, relationUserId, 1)
}

// 创建好友关系
func (rm RelationModel) CreateUserRelation(userId, relationUserId int, relationType int8) bool {
	relation := schema.Relation {
		UserId: userId,
		RelationUserId: relationUserId,
		RelationType: relationType,
	}

	relation.CreatedAt = time.Now()
	relation.UpdatedAt = time.Now()

	return rm.database.NewRecord(relation)
}

// 更新好友关系
func (rm RelationModel) UpdateUserRelation(relationId int, updatedInfo map[string]interface{}) {
	var relation schema.Relation

	rm.database.Where("id = ?", relationId).First(&relation)

	rm.database.Model(&relation).Update(updatedInfo)
}

// 删除好友关系
func (rm RelationModel) DeleteUserRelation(relationId int) {
	var relation schema.Relation

	rm.database.Where("id = ?", relationId).First(&relation)

	rm.database.Delete(&relation)
}
