package database

import (
  "../config"
  "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/mysql"
)

var fdatabase *gorm.DB

// custom the table name prefix 
func defaultTableNameHandler() {
  gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
    return "yd_" + defaultTableName
  }
}

func init() {
  defaultTableNameHandler()
}

func Connect(databaseType string) {
  database, err := gorm.Open(databaseType, config.DataBase[databaseType])

  if err != nil {
    panic(err)
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
