package model

import (
  "../database"
  "github.com/jinzhu/gorm"
)

type RelationModel struct {
  database        *gorm.DB
}

func NewRelationModel() RelationModel {
  return RelationModel { database.GetDataBase() }
}
