package controller

import (
	"yuedu/model"
	"yuedu/schema"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var articleInfo model.FullArticleInfo
	model.NewArticleModel().GetFullArticleInfo(0, &articleInfo)

	// - 文章作者信息
	var userInfo model.FullUserInfo
	model.NewUserModel().GetFullUserInfo(articleInfo.Anchor, &userInfo)

	var indexData map[string]interface{} = map[string]interface{} {
		"channels": channels,
		"article": articleInfo,
		"userInfo": userInfo,
	}

	return indexData
}

func (ic IndexController) GetArticleIno(c *gin.Context) {
	articleId, articleErr := strconv.ParseInt(c.Param("id"), 10, 0)

	if articleErr != nil {
		panic(articleErr)
	}

	var articleInfo model.ArticleInfo = model.NewArticleModel().GetArticleInfoById(int(articleId))

	c.JSON(200, articleInfo)
}
