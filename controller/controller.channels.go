package controller

import (
	"strconv"
	"yuedu/model"
	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	Controller
}

func (cc ChannelController) RenderChannel(c *gin.Context) {

	channelId, _ := strconv.Atoi(c.Param("channelId"))

	channelData := cc.GetChannelData(channelId)

	if userInfo, isExist := cc.GetUserInfo(c); isExist {
		channelData["userInfo"] = userInfo;
	}

	c.HTML(200, "wchannels.html", channelData)
}

func (cc ChannelController) GetChannelData(channelId int) map[string]interface{} {

	articleModel := model.NewArticleModel()

	return map[string]interface{} {
		"channels": model.NewChannelModel().GetAllChannels(),
		"articles": articleModel.GetArticlesByChannelId(channelId, 0, 10),
		"topArticles": articleModel.GetTopArticles(channelId, 10),
		"count": articleModel.GetArticleCountByChannelId(channelId),
	}
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

	articleModel := model.NewArticleModel()

	articles := articleModel.GetArticlesByChannelId(channelId, page, size)

	count := articleModel.GetArticleCountByChannelId(channelId)

	c.JSON(200, map[string]interface{} {"articles": articles, "count": count, "rc": "0"})

}