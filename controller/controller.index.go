package controller

import (
	"strconv"
	"yuedu/model"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	Controller
}

func (ic IndexController) HandleIndex(c *gin.Context) {
	indexData := ic.getIndexData()

	if userInfo, isExist := ic.GetUserInfo(c); isExist {
		indexData["userInfo"] = userInfo
	}

	c.HTML(200, "windex.html", indexData)
}

func (ic IndexController) getIndexData() map[string]interface{} {

	indexArticleId := 354

	channelModel, articleModel, commentModal := model.NewChannelModel(), model.NewArticleModel(), model.NewCommentlModel()

	relativeArticles := articleModel.GetReleasedArticlesByArticleId(indexArticleId, 20)

	return map[string]interface{} {
		"channels": channelModel.GetAllChannels(),
		"article": articleModel.GetArticleInfoById(indexArticleId),
		"relativeArticlesArray": ic.AdaptRelativeArticlesArray(relativeArticles),
		"relativeArticlesByOtherChannel": articleModel.GetOtherChannelLastArticlesByArticleId(indexArticleId),
		"comments": commentModal.GetArticleCommentInfos(indexArticleId),
	}
}

func (ic IndexController) AdaptRelativeArticlesArray(relativeArticles []model.SimpleArticleInfo) [][]model.SimpleArticleInfo {
	var relativeArticlesArray [][]model.SimpleArticleInfo

	for index, article := range relativeArticles {
		if index % 4 > 0 {
			relativeArticlesArray[len(relativeArticlesArray)-1] = append(relativeArticlesArray[len(relativeArticlesArray)-1], article)
		} else {
			relativeArticlesArray = append(relativeArticlesArray, []model.SimpleArticleInfo{article})
		}
	}

	return relativeArticlesArray
}

func (ic IndexController) HandleComment(c *gin.Context) {
	articleIdString, hasArticleIdString := c.GetQuery("articleId")

	if !hasArticleIdString {
		c.JSON(200, map[string]string{"rc": "1", "msg": "articleId 不能为空"})
		return
	}

	articleId, articleIdErr := strconv.Atoi(articleIdString)

	if articleIdErr != nil {
		c.JSON(200, map[string]string{"rc": "1", "msg": "articleId 转换失败"})
		return
	}

	var comments []model.CommentInfo = model.NewCommentlModel().GetArticleCommentInfos(int(articleId))

	c.JSON(200, map[string]interface{}{"rc": "0", "data": comments})
}
