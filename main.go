package main

import (
  "yuedu/router"
  "yuedu/database"
  "github.com/gin-gonic/gin"
)

var app *gin.Engine = gin.Default()

func init() {
  router.SetUpRouter(app)
  database.Connect("mysql")
}

func main() {
  app.Run(":8080")
}
