package router

import (
	"github.com/gin-gonic/gin"
	"yuedu/controller"
)

var loginController controller.LoginController

var loginRouter []descriptor = []descriptor {
	descriptor{
		path: "/captcha",
		method: "GET",
		handlers: []gin.HandlerFunc { loginController.GetCaptcha },
	},
	descriptor{
		path: "/register",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.HandleRegister },
	},
	descriptor{
		path: "/login",
		method: "POST",
		handlers: []gin.HandlerFunc { loginController.Login },
	},
}