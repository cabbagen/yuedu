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
