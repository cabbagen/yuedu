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

func (am ArticleModel) GetLastArticleByChannel(article *schema.Article, channelId int) {
  am.database.Where(map[string]interface{} {"channel_id": channelId}).Last(article)
}



