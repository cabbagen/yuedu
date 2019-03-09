package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
)

type CollectionModel struct {
	database        *gorm.DB
}

func NewCollectionModel() CollectionModel {
	return CollectionModel { database.GetDataBase() }
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
