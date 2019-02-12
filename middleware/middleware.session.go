package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

type SessionMiddleware struct {
	key      string
	name     string
}

func (sm SessionMiddleware) Handle() gin.HandlerFunc {
	store := cookie.NewStore([]byte(sm.key))

	return sessions.Sessions(sm.name, store)
}

func NewSessionMiddleware() SessionMiddleware {
	return SessionMiddleware{
		key: "secret",
		name: "session",
	}
}
