package model

import (
	"yuedu/schema"
	"yuedu/database"
	"github.com/jinzhu/gorm"
)

type ArticleModel struct {
	database        *gorm.DB
}

func NewArticleModel() ArticleModel {
	return ArticleModel { database.GetDataBase() }
}

type FullArticleInfo struct {
	schema.Article
	Supports            int    `json="supports"`
	Collections         int    `json="collections"`
}

func (am ArticleModel) GetFullArticleInfo(articleId int, fullArticleInfo *FullArticleInfo) {
	var articleInfo schema.Article

	if articleId > 0 {
		am.database.First(&articleInfo, articleId)
	} else {
		am.database.Where(map[string]interface{} {"channel_id": 1}).Last(&articleInfo)
	}

	(*fullArticleInfo).Article = articleInfo

	am.database.Table("yd_articles").
		Select("count(yd_collections.id) as collections, count(yd_supports.id) as supports").
		Where("yd_users.id = ?", articleInfo.ID).
		Joins("left join yd_collections on yd_collections.article_id = yd_articles.id").
		Joins("left join yd_supports on yd_supports.article_id = yd_articles.id").
		Group("yd_collections.article_id, yd_supports.article_id").
		Scan(fullArticleInfo)
}

