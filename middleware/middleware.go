package middleware

import "github.com/gin-gonic/gin"

type HandleMiddleware interface {
	Handle(c *gin.Context)
}

func SetUpMiddleware(app *gin.Engine) {
	app.Use(gin.Logger(), gin.Recovery())
}