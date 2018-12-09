package model

import (
	"yuedu/database"
	"yuedu/schema"
	"github.com/jinzhu/gorm"
)

type ChannelModel struct {
	database        *gorm.DB
}

func NewChannelModel() ChannelModel {
	return ChannelModel { database.GetDataBase() }
}

// - 获取所有频道
func (cm ChannelModel) FindAll(channles *[]schema.Channel) {
	cm.database.Find(channles)
}
