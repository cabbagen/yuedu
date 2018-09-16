package model

import (
  "../database"
  "github.com/jinzhu/gorm"
)

type SupportModel struct {
  database        *gorm.DB
}

func NewSupportModel() SupportModel {
  return SupportModel { database.GetDataBase() }
}
