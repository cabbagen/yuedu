package router

import (
  "../config"
  "github.com/gin-gonic/gin"
)

// custom route description for decoupling router modules
type descriptor struct {
  path        string
  method      string
  handlers    []gin.HandlerFunc
}

// the global routes for app engine
var routes []descriptor

// load all router in router modules into global router
func init() {
  routes = append(routes, indexRouter...)
}

func SetUpRouter(engine *gin.Engine) {
  registRoute(engine)
  bindRouteTmpl(engine)
}

func registRoute(engine *gin.Engine) {
  for _, route := range routes {
    engine.Handle(route.method, route.path, route.handlers...)
  }
  if config.Application["static"] != "" {
    engine.Static("static", config.Application["static"])
  }
}

func bindRouteTmpl(engine *gin.Engine) {
  engine.LoadHTMLGlob("./views/**/*.html")
}
