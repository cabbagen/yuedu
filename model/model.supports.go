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

	result := sm.database.Table("yd_supports").Where("state = 1 and article_id = ?", articleId).Count(&supportCount)

	if result.Error != nil {
		return supportCount, result.Error
	}

	return supportCount, nil
}

// 获取评论点赞数目
func (sm SupportModel) GetSupportCountForComment(commentId int) (int, error) {
	var supportCount int = 0

	if result := sm.database.Table("yd_supports").Where("state = 1 and comment_id = ?", commentId).Count(&supportCount); result.Error != nil {
		return supportCount, result.Error
	}

	return supportCount, nil
}

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

func (sm SupportModel) DeleteSupportArticle(userId, articleId int) error {
	result := sm.database.Table("yd_supports").
		Where("state = 1 and type = 1 and user_id = ? and article_id = ?", userId, articleId).
		Delete(&schema.Support{})

	return result.Error
}
