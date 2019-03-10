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

// 查询用户的关注粉丝列表
func (rm RelationModel) GetUserFollowers(userId int) ([]schema.User, error) {
	var singleFollowers []schema.User

	singleResult := rm.database.Table("yd_relations").
		Select("yd_users.*").
		Where("relation_user_id = ? and relation_type = 1", userId).
		Joins("inner join yd_users on yd_users.id = yd_relations.user_id").
		Find(&singleFollowers)

	if singleResult.Error != nil {
		return singleFollowers, singleResult.Error
	}

	commonFollowers, error := rm.GetRelationUsers(userId)

	if error != nil {
		return commonFollowers, error
	}

	singleFollowers = append(singleFollowers, commonFollowers...)

	return singleFollowers, nil
}

// 获取相互关注的用户
func (rm RelationModel) GetRelationUsers(userId int) ([]schema.User, error) {
	var relations []schema.Relation

	result := rm.database.Table("yd_relations").
		Select("*").
		Where("(user_id = ? or relation_user_id = ?) and relation_type = 2", userId, userId).
		Find(&relations)

	if result.Error != nil {
		return []schema.User{}, result.Error
	}

	var relatedUserIds []int

	for _, relation := range relations {
		if relation.UserId == userId {
			relatedUserIds = append(relatedUserIds, relation.RelationUserId)
		} else {
			relatedUserIds = append(relatedUserIds, relation.UserId)
		}
	}

	users, error := NewUserModel().GetUserInfoByUserIds(relatedUserIds)

	if error != nil {
		return users, error
	}

	return users, nil
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
	var singleFollowings []schema.User

	singleResult := rm.database.Table("yd_relations").
		Select("yd_users.*").
		Where("user_id = ? and relation_type = 1", userId).
		Joins("inner join yd_users on yd_users.id = yd_relations.relation_user_id").
		Find(&singleFollowings)

	if singleResult.Error != nil {
		return singleFollowings, singleResult.Error
	}


	commonFollowings, error := rm.GetRelationUsers(userId)

	if error != nil {
		return commonFollowings, error
	}

	singleFollowings = append(singleFollowings, commonFollowings...)

	return singleFollowings, nil
}

// 查询用户正在关注的人员数量
func (rm RelationModel) GetUserFollowingCount(userId int) (int, error) {
	var count int

	var condition string = "( (user_id = ? or relation_user_id = ?) and relation_type = 2 ) or ( user_id = ? and relation_type = 1 )"

	if result := rm.database.Table("yd_relations").Where(condition, userId, userId, userId).Count(&count); result.Error != nil {
		return count, result.Error
	}

	return count, nil
}

// 添加关注
func (rm RelationModel) CreateUserFollowing(userId, relatedUserId int) error {
	var relation schema.Relation

	var condition string = "user_id = ? and relation_user_id = ? and relation_type = 1"

	// 查询该用户是否被关注过
	// 如果被关注过，则修改这条记录
	// 如果未被关注过，则新建一条记录
	if result := rm.database.Table("yd_relations").Where(condition, relatedUserId, userId).First(&relation); result.Error != nil && !result.RecordNotFound() {
		return result.Error
	}

	if relation.ID > 0 {
		if error := rm.UpdateUserRelation(relation.ID, map[string]interface{} {"relation_type": 2}); error != nil {
			return error
		}
		return nil
	}

	if error := rm.CreateUserRelation(userId, relatedUserId, 1); error != nil {
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

// 获取两个用户之间的关系
func (rm RelationModel) GetUsersRelations(userId, relatedUserId int) (int, error) {
	var relation schema.Relation
	var followCondition = "relation_type = 1 and user_id = ? and relation_user_id = ?"

	followingResult := rm.database.Table("yd_relations").Where(followCondition, userId, relatedUserId).First(&relation)

	if followingResult.Error != nil && !followingResult.RecordNotFound() {
		return 0, followingResult.Error
	}

	// userId  关注  relatedUserId
	if relation.ID > 0 {
		return 1, nil
	}

	followerResult := rm.database.Table("yd_relations").Where(followCondition, relatedUserId, userId).First(&relation)

	if followerResult.Error != nil && !followerResult.RecordNotFound() {
		return 0, followerResult.Error
	}

	// relatedUserId  关注  userId
	if relation.ID > 0 {
		return 2, nil
	}

	var commonCondition = "relation_type = 2 and ((user_id = ? and relation_user_id = ?) or (user_id = ? and relation_user_id = ?))"

	commonResult := rm.database.Table("yd_relations").
		Where(commonCondition, userId, relatedUserId, relatedUserId, userId).
		First(&relation)

	if commonResult.Error != nil && !commonResult.RecordNotFound() {
		return 0, commonResult.Error
	}

	// userId  relatedUserId 相互关注
	if relation.ID > 0 {
		return 3, nil
	}

	return 0, nil
}

// 取消关注
func (rm RelationModel) CancelUserFollowing(userId, relatedUserId int) error {
	var relation schema.Relation

	// 如果存在单向的关系
	// 则直接删除该记录
	singleResult := rm.database.Table("yd_relations").
		Where("user_id = ? and relation_user_id = ? and relation_type = 1", userId, relatedUserId).
		First(&relation)
	
	if singleResult.Error != nil && !singleResult.RecordNotFound() {
		return singleResult.Error
	}
	
	if relation.ID > 0 {
		return rm.DeleteUserRelation(relation.ID)
	}
	
	// 如果两个用户已经相互关注
	// 则需要修改这条记录
	condition := "relation_type = 2 and ((user_id = ? and relation_user_id = ?) or (user_id = ? and relation_user_id = ?))"
	
	commonResult := rm.database.Table("yd_relations").
		Where(condition, userId, relatedUserId, relatedUserId, userId).
		First(&relation)
	
	if commonResult.Error != nil && !commonResult.RecordNotFound() {
		return commonResult.Error
	}

	updatedInfo := map[string]interface{} {
		"user_id": relatedUserId,
		"relation_user_id": userId,
		"relation_type": 1,
	}
	if relation.ID > 0 {
		return rm.UpdateUserRelation(relation.ID, updatedInfo)
	}

	return nil
}
