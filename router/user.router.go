package router

import (
	"yuedu/controller"
	"github.com/gin-gonic/gin"
)

var userController controller.UserController

var userRouter []descriptor = []descriptor {
	descriptor{
		path: "/user/:userId",
		method: "GET",
		handlers: []gin.HandlerFunc { userController.RenderUser },
	},
	//descriptor{
	//	path: "/users/:userId",
	//	method: "GET",
	//	handlers: []gin.HandlerFunc { userController.GetUserSampleInfo },
	//},
	descriptor{
		path: "/followers",
		method: "GET",
		handlers: []gin.HandlerFunc { userController.HandleFollows },
	},
	descriptor{
		path: "/followings",
		method: "GET",
		handlers: []gin.HandlerFunc { userController.HandleFollowings },
	},
}
