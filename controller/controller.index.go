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
	indexData, error := ic.getIndexData()

	if error != nil {
		ic.HandleThrowException(c, error)
	}

	if userInfo, isExist := ic.GetUserInfo(c); isExist {
		indexData["userInfo"] = userInfo
	}

	c.HTML(200, "windex.html", indexData)
}

func (ic IndexController) getIndexData() (map[string]interface{}, error) {

	var indexData map[string]interface{} = make(map[string]interface{})

	indexArticleId := 354

	channelModel, articleModel, commentModal := model.NewChannelModel(), model.NewArticleModel(), model.NewCommentlModel()

	relativeArticles, error := articleModel.GetReleasedArticlesByArticleId(indexArticleId, 20)

	if error != nil {
		return indexData, error
	}

	if channels, error := channelModel.GetAllChannels(); error != nil {
		return indexData, error
	} else {
		indexData["channels"] = channels
	}

	if article, error := articleModel.GetArticleInfoById(indexArticleId); error != nil {
		return indexData, error
	} else {
		indexData["article"] = article
	}

	if relativeArticlesByOtherChannel, error := articleModel.GetOtherChannelLastArticlesByArticleId(indexArticleId); error != nil {
		return indexData, error
	} else {
		indexData["relativeArticlesByOtherChannel"] = relativeArticlesByOtherChannel
	}

	if comments, error := commentModal.GetArticleCommentInfos(indexArticleId); error != nil {
		return indexData, error
	} else {
		indexData["comments"] = comments
	}

	indexData["relativeArticlesArray"] = ic.AdaptRelativeArticlesArray(relativeArticles)

	return indexData, nil
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

	articleIdString, isExist := c.GetQuery("articleId")

	if !isExist {
		c.JSON(200, map[string]string{"rc": "1", "msg": "articleId 不能为空"})
		return
	}

	articleId, error := strconv.Atoi(articleIdString)

	if error != nil {
		c.JSON(200, map[string]string{"rc": "1", "msg": "articleId 转换失败"})
		return
	}

	comments, error := model.NewCommentlModel().GetArticleCommentInfos(int(articleId))

	if error != nil {
		c.JSON(200, map[string]string{"rc": "2", "msg": "数据库操作失败"})
		return
	}

	c.JSON(200, map[string]interface{}{"rc": "0", "data": comments})
}
