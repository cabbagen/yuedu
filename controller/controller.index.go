package controller

import (
	"yuedu/model"
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

	var indexData map[string]interface{} = make(map[string]interface{})

	var indexArticleId int = 354

	var channelModel model.ChannelModel = model.NewChannelModel()

	var articleModel model.ArticleModel = model.NewArticleModel()

	var commentModal model.CommentlModel = model.NewCommentlModel()

	var relativeArticles []model.SimpleArticleInfo = articleModel.GetReleasedArticlesByArticleId(indexArticleId, 20)

	indexData["channels"] = channelModel.GetAllChannels()

	indexData["article"] = articleModel.GetArticleInfoById(indexArticleId)

	indexData["relativeArticlesArray"] = ic.AdaptRelativeArticlesArray(relativeArticles)

	indexData["relativeArticlesByOtherChannel"] = articleModel.GetOtherChannelLastArticlesByArticleId(indexArticleId)

	indexData["comments"] = commentModal.GetArticleCommentInfos(indexArticleId)

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

func (ic IndexController) HandleComment(c *gin.Context) {
	articleIdString, hasArticleIdString := c.GetQuery("articleId")

	if !hasArticleIdString {
		c.JSON(200, map[string]string {"rc": "1", "msg": "articleId 不能为空"})
		return
	}

	articleId, articleIdErr := strconv.Atoi(articleIdString)

	if articleIdErr != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "articleId 转换失败"})
		return
	}

	var comments []model.CommentInfo = model.NewCommentlModel().GetArticleCommentInfos(int(articleId))

	c.JSON(200, map[string]interface{} {"rc": 0, "data": comments})
}
