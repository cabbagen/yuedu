package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenMiddleware struct {
	Key       string
	Name      string
}

type TokenClaims struct {
	UserInfo      string  `json:"userInfo"`
	jwt.StandardClaims
}

// 验证 token 中间件函数
func (tm TokenMiddleware) Handle(c *gin.Context) {
	tokenString := c.GetHeader("token")

	if tokenString == "" {
		c.AbortWithStatusJSON(200, map[string]string {"rc": "8", "msg": "当前没有权限查看"})
	}

	_, ok := tm.ValidateToken(tokenString)

	if !ok {
		c.AbortWithStatusJSON(200, map[string]string {"rc": "8", "msg": "token 解析错误"})
	}

	c.Next()
}

func (tm TokenMiddleware) SignToken(userInfo string) (string, error) {
	claims := TokenClaims{
		UserInfo: userInfo,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, error := token.SignedString([]byte(tm.Key))

	if error != nil {
		return "", error
	}

	return tokenString, nil
}

func (tm TokenMiddleware) ValidateToken(tokenString string) (string, bool) {
	token, error := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tm.Key), nil
	})

	if error != nil {
		return "", false
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.UserInfo, true
	}

	return "", false
}

func NewTokenMiddleware() TokenMiddleware {
	return TokenMiddleware { Key: "token-middleware", Name: "token" }
}


