package model

import (
  "yuedu/database"
  "github.com/jinzhu/gorm"
)

type ChannelModel struct {
  database        *gorm.DB
}

func NewChannelModel() ChannelModel {
  return ChannelModel { database.GetDataBase() }
}
