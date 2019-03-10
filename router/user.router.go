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
}
