package router

import (
  "os"
  "log"
  "../config"
  "path/filepath"
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

var TmplRootDir string = "./views/"

var tmplFiles []string

var tmplFileExt string = ".tmpl"

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
  tmplErr := filepath.Walk(TmplRootDir, func (path string, fileInfo os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if !fileInfo.IsDir() && filepath.Ext(path) == tmplFileExt {
      tmplFiles = append(tmplFiles, path)
    }
    return nil
  })

  if tmplErr != nil {
    log.Println("绑定模板出错：", tmplErr)
  }
  engine.LoadHTMLFiles(tmplFiles...)
}
