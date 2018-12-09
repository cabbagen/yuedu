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

var routes []descriptor

func init() {
	routes = append(routes, indexRouter...)
}

func SetUpRouter(engine *gin.Engine) {
	for _, route := range routes {
		engine.Handle(route.method, route.path, route.handlers...)
	}
	if config.Application["static"] != "" {
		engine.Static("static", config.Application["static"])
	}
}

func demo() {

}