package main

import (
	"github.com/gin-gonic/gin"
	"yuedu/ctemplate"
	"yuedu/database"
	"yuedu/router"
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
