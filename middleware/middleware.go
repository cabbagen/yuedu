package middleware

import "github.com/gin-gonic/gin"

type HandleMiddleware interface {
	Handle() gin.HandlerFunc
}

func SetUpMiddleware(engine *gin.Engine) {
	handles := []gin.HandlerFunc {
		NewSessionMiddleware().Handle(),
	}

	engine.Use(handles...)
}
