package database

import (
  "../config"
  "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/mysql"
)

var fdatabase *gorm.DB

func init() {
  defaultTableNameHandler()
}

func defaultTableNameHandler() {
  gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
    return "yd_" + defaultTableName
  }
}

func Connect(databaseType string) {
  database, connectErr := gorm.Open(databaseType, config.DataBase[databaseType])

  if connectErr != nil {
    panic(connectErr)
  }

  fdatabase = database;
}

func GetDataBase() *gorm.DB {
  if fdatabase == nil {
    panic("database is nil, you need invoke database.Connect function")
  }
  return fdatabase;
}

func Destory() bool {
  if fdatabase != nil {
    fdatabase = nil
  }
  return true
}
