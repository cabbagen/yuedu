package controller

import (
	"yuedu/model"
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

	var indexData map[string]interface{} = make(map[string]interface{})

	var indexArticleId int = 354

	var channelModel model.ChannelModel = model.NewChannelModel()

	var articleModel model.ArticleModel = model.NewArticleModel()

	var relativeArticles []model.SimpleArticleInfo = articleModel.GetReleasedArticlesByArticleId(indexArticleId, 20)

	indexData["channels"] = channelModel.GetAllChannels()

	indexData["article"] = articleModel.GetArticleInfoById(indexArticleId)

	indexData["relativeArticlesArray"] = ic.AdaptRelativeArticlesArray(relativeArticles)

	indexData["relativeArticlesByOtherChannel"] = articleModel.GetOtherChannelLastArticlesByArticleId(indexArticleId)

	return indexData
}

func (ic IndexController) AdaptRelativeArticlesArray(relativeArticles []model.SimpleArticleInfo) [][]model.SimpleArticleInfo {
	var relativeArticlesArray [][]model.SimpleArticleInfo

	for index, article := range relativeArticles {
		if index % 4 > 0 {
			relativeArticlesArray[len(relativeArticlesArray) - 1] = append(relativeArticlesArray[len(relativeArticlesArray) - 1], article)
		} else {
			relativeArticlesArray = append(relativeArticlesArray, []model.SimpleArticleInfo{article})
		}
	}

	return relativeArticlesArray
}
