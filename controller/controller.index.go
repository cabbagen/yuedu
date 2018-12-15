package controller

import (
	"yuedu/model"
	"yuedu/schema"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Controller
}


func (ic IndexController) HandleIndex(c *gin.Context) {
	var indexData map[string]interface{} = ic.getIndexData()

	c.HTML(200, "windex.html", indexData)
}

func (ic IndexController) getIndexData() map[string]interface{} {

	// - 获取所有频道
	var channels []schema.Channel
	model.NewChannelModel().FindAll(&channels)

	// - 获取类目下的文章
	var article schema.Article
	model.NewArticleModel().GetLastArticleByChannel(&article, 1)

	// - 文章作者信息
	var userInfo model.FullUserInfo
	model.NewUserModel().GetFullUserInfo(article.Anchor, &userInfo)

	var indexData map[string]interface{} = map[string]interface{} {
		"channels": channels,
		"article": article,
		"userInfo": userInfo,
	}

	return indexData
	
}
