package controller

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Controller
}

func (ic IndexController) HandleIndex(c *gin.Context) {
	c.HTML(200, "index.tmpl", map[string]string {"title": "hello golang"})
}
