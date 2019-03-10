package router

import (
	"yuedu/controller"
	"github.com/gin-gonic/gin"
)

var channelController controller.ChannelController

var channelRouter []descriptor = []descriptor {
	descriptor{
		path: "/channels/:channelId",
		method: "GET",
		handlers: []gin.HandlerFunc{ channelController.RenderChannel },
	},
	descriptor{
		path: "/channel/articles",
		method: "GET",
		handlers: []gin.HandlerFunc { channelController.GetChannelArticles },
	},
}
