package main

import (
  "./router"
  "./database"
  "github.com/gin-gonic/gin"
)

var app *gin.Engine = gin.Default()

func init() {
  router.SetUp(app)
  database.Connect("mysql")
}

func main() {
  app.Run(":8080")
}
