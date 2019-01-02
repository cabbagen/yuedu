package controller

import (
	"yuedu/model"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Controller
}


func (ic IndexController) HandleIndex(c *gin.Context) {
	var indexData map[string]interface{} = ic.getIndexData()

	c.HTML(200, "windex.html", indexData)

	//c.JSON(200, indexData)
}

func (ic IndexController) getIndexData() map[string]interface{} {

	var indexData map[string]interface{} = make(map[string]interface{})

	indexData["channels"] = model.NewChannelModel().GetAllChannels()

	indexData["article"] = model.NewArticleModel().GetArticleInfoById(354)

	return indexData
}
