package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
	"yuedu/schema"
)

type CollectionModel struct {
	database        *gorm.DB
}

func NewCollectionModel() CollectionModel {
	return CollectionModel { database.GetDataBase() }
}

// 获取文章的收藏数目
func (cm CollectionModel) GetCollectionCountForArticle(articleId int) (int, error) {
	var collectedCount int = 0

	result := cm.database.Table("yd_collections").Where("state = 1 and article_id = ?", articleId).Count(&collectedCount)

	if result.Error != nil {
		return collectedCount, result.Error
	}

	return collectedCount, nil
}

func (cm CollectionModel) GetUserCollectedArticles(userId int) ([]SmallArticleInfo, error) {
	var articles []SmallArticleInfo

	rows, error := cm.database.Table("yd_collections").
		Select("yd_articles.id, title, author, anchor, cover_img").
		Where("yd_collections.user_id = ?", userId).
		Joins("inner join yd_articles on yd_articles.id = yd_collections.article_id").
		Rows()

	if error != nil {
		return articles, error
	}

	for rows.Next() {
		article := SmallArticleInfo{}

		if error := rows.Scan(&article.Id, &article.Title, &article.Author, &article.AnchorName, &article.CoverImg); error != nil {
			return articles, error
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (cm CollectionModel) IsCollectedArticle(userId, articleId int) (bool, error) {
	var collection schema.Collection

	result := cm.database.Table("yd_collections").
		Where("user_id = ? and article_id = ? and state = 1", userId, articleId).
		First(&collection)

	if result.Error != nil && !result.RecordNotFound() {
		return false, result.Error
	}

	if collection.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (cm CollectionModel) CreateCollectArticle(userId, articleId int) error {
	collection := schema.Collection{
		State: 1,
		UserId: userId,
		ArticleId: articleId,
	}

	result := cm.database.Table("yd_collections").Create(&collection)

	return result.Error
}

func (cm CollectionModel) DeleteCollectArticle(userId, articleId int) error {
	result := cm.database.Table("yd_collections").
		Where("state = 1 and user_id = ? and article_id = ?", userId, articleId).
		Delete(&schema.Collection{})

	return result.Error
}


