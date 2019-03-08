package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
)

type SupportModel struct {
	database        *gorm.DB
}

func NewSupportModel() SupportModel {
	return SupportModel { database.GetDataBase() }
}

// 获取文章的点赞数目
func (sm SupportModel) getSupportCountForArticle(articleId int) (int, error) {
	var supportCount int = 0

	if result := sm.database.Table("yd_supports").Where("article_id = ?", articleId).Count(&supportCount); result.Error != nil {
		return supportCount, result.Error
	}

	return supportCount, nil
}


// 获取评论点赞数目
func (sm SupportModel) getSupportCountForComment(commentId int) (int, error) {
	var supportCount int = 0

	if result := sm.database.Table("yd_supports").Where("comment_id = ?", commentId).Count(&supportCount); result.Error != nil {
		return supportCount, result.Error
	}

	return supportCount, nil
}
