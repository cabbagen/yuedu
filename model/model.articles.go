package model

import (
  "../database"
  "github.com/jinzhu/gorm"
)

type ArticleModel struct {
  database        *gorm.DB
}

func NewArticleModel() ArticleModel {
  return ArticleModel { database.GetDataBase() }
}
