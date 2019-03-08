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
func (cm ChannelModel) GetAllChannels() ([]schema.Channel, error) {
	var channels []schema.Channel

	if result := cm.database.Find(&channels); result.Error != nil {
		return channels, result.Error
	}

	return channels, nil
}
