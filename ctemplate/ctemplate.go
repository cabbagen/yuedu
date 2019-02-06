package ctemplate

import (
	"html/template"
	"github.com/gin-gonic/gin"
)

var templatePath string = "./views/**/*.html"

func SetUpTemplater(engine *gin.Engine) {
	var myFuncMap template.FuncMap = template.FuncMap {
		"unescaped": unescaped,
		"formatTimeDuring": formatTimeDuring,
		"formatTimeString": formatTimeString,
		"removeHtmlTags": removeHtmlTags,
	}
	engine.SetFuncMap(myFuncMap)
	engine.LoadHTMLGlob(templatePath)
}
