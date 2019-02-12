package middleware

import "github.com/gin-gonic/gin"

type HandleMiddleware interface {
	Handle(c *gin.Context)
}
