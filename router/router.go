package router

import (
	"yuedu/config"
	"github.com/gin-gonic/gin"
)

type descriptor struct {
	path        string
	method      string
	handlers    []gin.HandlerFunc
}

var routers []descriptor

func init() {
	routers = append(routers, indexRouter...)
	routers = append(routers, loginRouter...)
	routers = append(routers, channelRouter...)
	routers = append(routers, userRouter...)
}

func SetUpRouter(engine *gin.Engine) {
	for _, router := range routers {
		engine.Handle(router.method, router.path, router.handlers...)
	}
	if config.Application["static"] != "" {
		engine.Static("static", config.Application["static"])
	}
}
