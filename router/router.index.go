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
	descriptor{
		path: "/article/comment",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.HandleArticleComment },
	},
	descriptor{
		path: "/comment/comment",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.HandleCommentComment },
	},
	descriptor{
		path: "/comment/delete",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.HandleDeleteComment },
	},
	descriptor{
		path: "/comment/support",
		method: "POST",
		handlers: []gin.HandlerFunc { tokenMiddleware.Handle, indexController.HandleCommentSupport },
	},
}
