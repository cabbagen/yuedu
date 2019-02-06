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

	channelId, _ := strconv.Atoi(c.Param("channelId"))

	channelData := cc.GetChannelData(channelId)

	c.HTML(200, "wchannels.html", channelData)
}

func (cc ChannelController) GetChannelData(channelId int) map[string]interface{} {
	var channelData map[string]interface{} = make(map[string]interface{})
	var articleModel model.ArticleModel = model.NewArticleModel()

	channelData["channels"] = model.NewChannelModel().GetAllChannels()
	channelData["articles"] = articleModel.GetArticlesByChannelId(channelId, 0, 10)
	channelData["topArticles"] = articleModel.GetTopArticles(channelId, 10)
	channelData["count"] = articleModel.GetArticleCountByChannelId(channelId)


	return channelData
}

func (cc ChannelController) GetChannelArticles(c *gin.Context) {

	channelId, error := strconv.Atoi(c.DefaultQuery("channelId", "1"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	page, error := strconv.Atoi(c.DefaultQuery("page", "0"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	size, error := strconv.Atoi(c.DefaultQuery("size", "10"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	var articleModel model.ArticleModel = model.NewArticleModel()

	var articles []model.SimpleArticleInfo = articleModel.GetArticlesByChannelId(channelId, page, size)

	var count int = articleModel.GetArticleCountByChannelId(channelId)

	c.JSON(200, map[string]interface{} {"articles": articles, "count": count, "rc": "0"})

}