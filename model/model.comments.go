package model

import (
  "yuedu/database"
  "github.com/jinzhu/gorm"
)

type CommentlModel struct {
  database        *gorm.DB
}

func NewCommentlModel() CommentlModel {
  return CommentlModel { database.GetDataBase() }
}
