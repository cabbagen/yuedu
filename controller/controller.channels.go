package controller

import "github.com/gin-gonic/gin"

type ChannelController struct {
	Controller
}

func (cc ChannelController) RenderChannel(c *gin.Context) {

}

func (cc ChannelController) GetChannels(c *gin.Context) {
	channelId, hasChannelId := c.GetQuery("channelId")

	pageNo, hasPageNo := c.GetQuery("pageNo")

	pageSize, hasPageSize := c.GetQuery("pageSize")

	if !hasChannelId {
		c.JSON(200, "channelId 不能为空")
		return
	}

	if !hasPageNo {
		c.JSON(200, "pageNo 不能为空")
		return
	}

	if !hasPageSize {
		c.JSON(200, "pageSize 不能为空")
		return
	}


	c.JSON(200, "yeah!" + channelId + pageNo + pageSize)
}