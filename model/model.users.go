package model

import (
  "yuedu/database"
  "yuedu/schema"
  "github.com/jinzhu/gorm"
)

type UserModel struct {
  database        *gorm.DB
}

func NewUserModel() UserModel {
  return UserModel { database.GetDataBase() }
}

func (u UserModel) Query(users *[]schema.User) {
  u.database.Find(users)
}
