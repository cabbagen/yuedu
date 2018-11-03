package model

import (
  "yuedu/database"
  "github.com/jinzhu/gorm"
)

type AddressModel struct {
  database        *gorm.DB
}

func NewAddressModel() AddressModel {
  return AddressModel { database.GetDataBase() }
}
