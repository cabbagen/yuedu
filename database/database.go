package database

import (
  "../config"
  "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func init() {
  defaultTableNameHandler()
}

func defaultTableNameHandler() {
  gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
    return "cb_" + defaultTableName
  }
}

func Connect() {
  db, err := gorm.Open("mysql", config.DataBase["mysql"])

  // defer db.Close()

  if err != nil {
    panic(err)
  }

  database = db
}

func GetDataBase() *gorm.DB {
  return database;
}
