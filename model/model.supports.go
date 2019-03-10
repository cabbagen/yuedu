package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"yuedu/schema"
)

type SupportModel struct {
	database        *gorm.DB
}

func NewSupportModel() SupportModel {
	return SupportModel { database.GetDataBase() }
}

// 获取文章的点赞数目
func (sm SupportModel) GetSupportCountForArticle(articleId int) (int, error) {
	var supportCount int = 0

	result := sm.database.Table("yd_supports").Where("type = 1 and state = 1 and article_id = ?", articleId).Count(&supportCount)

	if result.Error != nil {
		return supportCount, result.Error
	}

	return supportCount, nil
}

// 获取评论点赞数目
func (sm SupportModel) GetSupportCountForComment(commentId int) (int, error) {
	var supportCount int = 0

	if result := sm.database.Table("yd_supports").Where("type = 2 and state = 1 and comment_id = ?", commentId).Count(&supportCount); result.Error != nil {
		return supportCount, result.Error
	}

	return supportCount, nil
}

// 是否点赞过该文章
func (sm SupportModel) IsSupportArticle(userId, articleId int) (bool, error) {
	var support schema.Support

	result := sm.database.Table("yd_supports").
		Where("state = 1 and type = 1 and article_id = ? and user_id = ?", articleId, userId).
		First(&support)

	if result.Error != nil && !result.RecordNotFound() {
		return false, result.Error
	}

	if support.ID == 0 {
		return false, nil
	}

	return true, nil
}

// 是否点赞过该评论
func (sm SupportModel) IsSupportComment(userId, commentId int) (bool, error) {
	var support schema.Support

	result := sm.database.Table("yd_supports").
		Where("state = 1 and type = 2 and comment_id = ? and user_id = ?", commentId, userId).
		First(&support)

	if result.Error != nil && !result.RecordNotFound() {
		return false, result.Error
	}

	if support.ID == 0 {
		return false, nil
	}

	return true, nil
}

// 文章点赞
func (sm SupportModel) CreateSupportArticle(userId, articleId int) error {
	support := schema.Support{
		Type: 1,
		State: 1,
		UserId: userId,
		ArticleId: articleId,
		CommentId: 0,
	}

	result := sm.database.Table("yd_supports").Create(&support)

	return result.Error
}

// 评论点赞
func (sm SupportModel) CreateSupportComment(userId, commentId int) error {
	support := schema.Support{
		Type: 2,
		State: 1,
		UserId: userId,
		ArticleId: 0,
		CommentId: commentId,
	}

	result := sm.database.Table("yd_supports").Create(&support)

	return result.Error
}

// 文章取消点赞
func (sm SupportModel) DeleteSupportArticle(userId, articleId int) error {
	result := sm.database.Table("yd_supports").
		Where("state = 1 and type = 1 and user_id = ? and article_id = ?", userId, articleId).
		Delete(&schema.Support{})

	return result.Error
}

// 评论取消点赞
func (sm SupportModel) DeleteSupportComment(userId, commentId int) error {
	result := sm.database.Table("yd_supports").
		Where("state = 1 and type = 2 and user_id = ? and comment_id = ?", userId, commentId).
		Delete(&schema.Support{})

	return result.Error
}

// 删除评论
func (sm SupportModel) DeleteSupportCommentByCommentIds(commentIds []int) error {
	result := sm.database.Table("yd_supports").
		Where("state = 1 and type = 2 and comment_id in (?)", commentIds).
		Delete(&schema.Support{})

	return result.Error
}