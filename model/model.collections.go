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
