package router

import (
	"yuedu/controller"
	"github.com/gin-gonic/gin"
	"yuedu/middleware"
)

var indexController controller.IndexController
var tokenMiddleware middleware.TokenMiddleware = middleware.NewTokenMiddleware()

var indexRouter []descriptor = []descriptor {
	descriptor{
		path: "/index",
		method: "GET",
		handlers:  []gin.HandlerFunc{ indexController.RenderIndex },
	},
	descriptor{
		path: "/article/:articleId",
		method: "GET",
		handlers: []gin.HandlerFunc { indexController.RenderArticle },
	},
	descriptor{
		path: "/comments",
		method: "GET",
		handlers: []gin.HandlerFunc { indexController.HandleComment },
	},
	descriptor{
		path: "/article/support",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.SupportArticle },
	},
	descriptor{
		path: "/article/collect",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.CollectArticle },
	},
	descriptor{
		path: "/following/anchor",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.HandleFollowingAnchor },
	},
}
