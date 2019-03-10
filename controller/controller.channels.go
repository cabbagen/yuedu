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

	channelId, error := strconv.Atoi(c.Param("channelId"))

	if error != nil {
		cc.HandleThrowException(c, error)
	}

	channelData, error := cc.GetChannelData(channelId)

	if error != nil {
		cc.HandleThrowException(c, error)
	}

	if userInfo, isExist := cc.GetUserInfo(c); isExist {
		channelData["userInfo"] = userInfo;
	}

	c.HTML(200, "wchannels.html", channelData)
}

func (cc ChannelController) GetChannelData(channelId int) (map[string]interface{}, error) {
	var channelData map[string]interface{} = make(map[string]interface{})

	articleModel := model.NewArticleModel()

	if channels, error := model.NewChannelModel().GetAllChannels(); error != nil {
		return channelData, error
	} else {
		channelData["channels"] = channels
	}

	if articles, error := articleModel.GetArticlesByChannelId(channelId, 0, 10); error != nil {
		return channelData, error
	} else {
		channelData["articles"] = articles
	}

	if topArticles, error := articleModel.GetTopArticles(channelId, 10); error != nil {
		return channelData, error
	} else {
		channelData["topArticles"] = topArticles
	}

	if count, error := articleModel.GetArticleCountByChannelId(channelId); error != nil {
		return channelData, error
	} else {
		channelData["count"] = count
	}

	return channelData, nil
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

	articles, error := articleModel.GetArticlesByChannelId(channelId, page, size)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "2", "msg": error.Error()})
		return
	}

	count, error := articleModel.GetArticleCountByChannelId(channelId)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "2", "msg": error.Error()})
		return
	}

	c.JSON(200, map[string]interface{} {"articles": articles, "count": count, "rc": "0"})
}
