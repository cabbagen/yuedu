package main

import (
  "./router"
  "github.com/gin-gonic/gin"
)

var app *gin.Engine = gin.Default()

func init() {
  router.SetUp(app)
}

func main() {
  app.Run(":8080")
}
