package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yuedu/model"
)

type ChannelController struct {
	Controller
}

func (cc ChannelController) RenderChannel(c *gin.Context) {

}

func (cc ChannelController) GetChannelArticles(c *gin.Context) {

	channelId, error := strconv.Atoi(c.Param("channelId"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	page, error := strconv.Atoi(c.DefaultQuery("page", "0"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	size, error := strconv.Atoi(c.DefaultQuery("size", "20"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	var articleModel model.ArticleModel = model.NewArticleModel()

	var articles []model.SimpleArticleInfo = articleModel.GetArticlesByChannelId(channelId, page, size)

	var count int = articleModel.GetArticleCountByChannelId(channelId)

	c.JSON(200, map[string]interface{} {"articles": articles, "count": count, "rc": "0"})

}