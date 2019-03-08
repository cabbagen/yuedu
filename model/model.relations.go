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
func (rm RelationModel) GetUserFollowers(userId int) ([]schema.User, error) {
	var followers []schema.User

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( relation_user_id = ? and relation_type = 1 )"

	result := rm.database.Table("yd_relations").
		Select("yd_users.*").
		Where(condition, userId, userId, userId).
		Joins("inner join yd_users on yd_users.id = yd_relations.relation_user_id").
		Find(&followers)

	if result.Error != nil {
		return followers, result.Error
	}

	return followers, nil
}

// 查询粉丝的数量
func (rm RelationModel) GetUserFollowerCount(userId int) (int, error) {
	var count int
	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( relation_user_id = ? and relation_type = 1 )"

	if result := rm.database.Table("yd_relations").Where(condition, userId, userId, userId).Count(&count); result.Error != nil {
		return count, result.Error
	}

	return count, nil
}

// 查询用户正在关注的人员的列表
func (rm RelationModel) GetUserFollowings(userId int) ([]schema.User, error) {
	var followings []schema.User

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( user_id = ? and relation_type = 1 )"

	// 互相关注
	result := rm.database.Table("yd_relations").
		Select("yd_users.*").
		Where(condition, userId, userId, userId).
		Joins("inner join yd_users on yd_users.id = yd_relations.relation_user_id").
		Find(&followings)

	if result.Error != nil {
		return followings, result.Error
	}

	return followings, nil
}

func (rm RelationModel) GetUserFollowingCount(userId int) (int, error) {
	var count int

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( user_id = ? and relation_type = 1 )"

	if result := rm.database.Table("yd_relations").Where(condition, userId, userId, userId).Count(&count); result.Error != nil {
		return count, result.Error
	}

	return count, nil
}

// 添加关注
func (rm RelationModel) CreateUserFollowing(userId, relationUserId int) error {
	var relation schema.Relation

	var condition string = "user_id = ? and relation_user_id = ? and relation_type = 1"

	// 查询该用户是否被关注过
	// 如果被关注过，则修改这条记录
	// 如果未被关注过，则新建一条记录
	if result := rm.database.Table("yd_relations").Where(condition, relationUserId, userId).First(&relation); result.Error != nil {
		return result.Error
	}

	if relation.ID > 0 {
		if error := rm.UpdateUserRelation(relation.ID, map[string]interface{} {"relation_type": 2}); error != nil {
			return error
		}
		return nil
	}

	if error := rm.CreateUserRelation(userId, relationUserId, 1); error != nil {
		return error
	}

	return nil
}

// 创建好友关系
func (rm RelationModel) CreateUserRelation(userId, relationUserId int, relationType int8) error {
	relation := schema.Relation {
		UserId: userId,
		RelationUserId: relationUserId,
		RelationType: relationType,
	}

	relation.CreatedAt = time.Now()
	relation.UpdatedAt = time.Now()

	if result := rm.database.Create(&relation); result.Error != nil {
		return result.Error
	}

	return nil
}

// 更新好友关系
func (rm RelationModel) UpdateUserRelation(relationId int, updatedInfo map[string]interface{}) error {
	if result := rm.database.Table("yd_relations").Where("id = ?", relationId).Update(updatedInfo); result.Error != nil {
		return result.Error
	}

	return nil
}

// 删除好友关系
func (rm RelationModel) DeleteUserRelation(relationId int) error {
	if result := rm.database.Where("id = ?", relationId).Delete(&schema.Relation{}); result.Error != nil {
		return result.Error
	}
	return nil
}
