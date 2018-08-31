package router

import (
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
	reduce(indexRouter)
}

func reduce(r []descriptor) {
	for _, route := range r {
		routes = append(routes, route)
	}
}

func registe(engine *gin.Engine) {
	for _, route := range routes {
		engine.Handle(route.method, route.path, route.handlers...)
	}
}

var TmplRootDir string = "./views/"

var tmplFiles []string = []string {
	TmplRootDir + "index.tmpl",
}

func bindTmpl(engine *gin.Engine) {
	engine.LoadHTMLFiles(tmplFiles...)
}

func SetUp(engine *gin.Engine) {
	registe(engine)
	bindTmpl(engine)
}
