package middleware

import "github.com/gin-gonic/gin"

type TokenMiddleware struct {
	key       string
	name      string
}

func (tm TokenMiddleware) Handle() gin.HandlerFunc {

}


