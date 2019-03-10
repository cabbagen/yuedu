package router

import (
	"yuedu/controller"
	"github.com/gin-gonic/gin"
)

var searchController controller.SearchController

var searchRouter []descriptor = []descriptor {
	descriptor{
		path: "/search",
		method: "GET",
		handlers: []gin.HandlerFunc{ searchController.RenderSearch },
	},
	descriptor{
		path: "/search/articles",
		method: "GET",
		handlers: []gin.HandlerFunc { searchController.GetSearchArticles },
	},
}
