package model

import (
	"yuedu/database"
	"github.com/jinzhu/gorm"
)

type ReplyModel struct {
	database        *gorm.DB
}

func NewReplyModel() ReplyModel {
	return ReplyModel { database.GetDataBase() }
}
