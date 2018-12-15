package main

import (
	"yuedu/router"
	"yuedu/ctemplate"
	"yuedu/database"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine = gin.Default()

func init() {
	database.Connect("mysql")

	router.SetUpRouter(app)

	ctemplate.SetUpTemplater(app)
}

func main() {
	app.Run(":8080")
}
