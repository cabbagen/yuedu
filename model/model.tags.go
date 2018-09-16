package model

import (
  "../database"
  "github.com/jinzhu/gorm"
)

type TagModel struct {
  database        *gorm.DB
}

func NewTagModel() TagModel {
  return TagModel { database.GetDataBase() }
}
