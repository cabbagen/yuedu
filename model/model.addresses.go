package model

import (
  "../database"
  "github.com/jinzhu/gorm"
)

type AddressModel struct {
  database        *gorm.DB
}

func NewAddressModel() AddressModel {
  return AddressModel { database.GetDataBase() }
}
