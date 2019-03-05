package router

import (
	"yuedu/controller"
	"github.com/gin-gonic/gin"
)

var userController controller.UserController

var userRouter []descriptor = []descriptor {
	descriptor{
		path: "/followers",
		method: "GET",
		handlers: []gin.HandlerFunc { userController.HandleFollows },
	},
	descriptor{
		path: "followings",
		method: "GET",
		handlers: []gin.HandlerFunc { userController.HanleFollowings },
	},
}
