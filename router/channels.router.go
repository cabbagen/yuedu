package router

import (
	"github.com/gin-gonic/gin"
	"yuedu/controller"
)

var channelController controller.ChannelController

var channelRouter []descriptor = []descriptor {
	descriptor{
		path: "/channels/:channelId",
		method: "GET",
		handlers:  []gin.HandlerFunc{ channelController.GetChannels },
	},
}
