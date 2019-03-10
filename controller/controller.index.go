package controller

import (
	"strconv"
	"yuedu/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type IndexController struct {
	Controller
}

func (ic IndexController) RenderIndex(c *gin.Context) {
	indexData, error := ic.GetIndexData(354)

	if error != nil {
		ic.HandleThrowException(c, error)
	}

	if userInfo, isExist := ic.GetUserInfo(c); isExist {
		indexData["userInfo"] = userInfo

		if extraInfo, error := ic.GetUserExtraInfo(userInfo.ID, 354); error == nil {
			indexData["extraInfo"] = extraInfo
		}
	}

	c.HTML(200, "windex.html", indexData)
}

func (ic IndexController) GetIndexData(articleId int) (map[string]interface{}, error) {

	var indexData map[string]interface{} = make(map[string]interface{})

	channelModel, articleModel, commentModal := model.NewChannelModel(), model.NewArticleModel(), model.NewCommentlModel()

	relativeArticles, error := articleModel.GetReleasedArticlesByArticleId(articleId, 20)

	if error != nil {
		return indexData, error
	}

	if channels, error := channelModel.GetAllChannels(); error != nil {
		return indexData, error
	} else {
		indexData["channels"] = channels
	}

	if article, error := articleModel.GetArticleInfoById(articleId); error != nil {
		return indexData, error
	} else {
		indexData["article"] = article
	}

	if relativeArticlesByOtherChannel, error := articleModel.GetOtherChannelLastArticlesByArticleId(articleId); error != nil {
		return indexData, error
	} else {
		indexData["relativeArticlesByOtherChannel"] = relativeArticlesByOtherChannel
	}

	if comments, error := commentModal.GetArticleCommentInfos(articleId); error != nil {
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

func (ic IndexController) GetUserExtraInfo(userId, articleId int) (map[string]interface{}, error) {
	extraInfo, error := model.NewUserModel().GetUserExtraInfo(userId, articleId)

	if error != nil {
		return map[string]interface{}{}, error
	}

	return extraInfo, nil
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

func (ic IndexController) RenderArticle(c *gin.Context) {
	articleId, error := strconv.Atoi(c.Param("articleId"))

	if error != nil {
		ic.HandleThrowException(c, error)
	}

	detailData, error := ic.GetIndexData(articleId)

	if error != nil {
		ic.HandleThrowException(c, error)
	}

	if userInfo, isExist := ic.GetUserInfo(c); isExist {
		detailData["userInfo"] = userInfo
		if extraInfo, error := ic.GetUserExtraInfo(userInfo.ID, articleId); error == nil {
			detailData["extraInfo"] = extraInfo
		}
	}

	c.HTML(200, "windex.html", detailData)
}


// 文章点赞
func (ic IndexController) SupportArticle(c *gin.Context) {
	articleIdString, isExist := c.GetPostForm("articleId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 articleId"))
	}

	articleId, error := strconv.Atoi(articleIdString)

	if error != nil {
		ic.HandleThrowException(c, errors.New("请传入 articleId"))
	}

	isSupported, isExist := c.GetPostForm("isSupported")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 isSupported"))
	}

	userInfo, _ := ic.GetUserInfo(c)

	if isSupported == "true" {
		if error := model.NewSupportModel().CreateSupportArticle(userInfo.ID, articleId); error != nil {
			c.JSON(200, map[string]string {"rc": "1", "msg": error.Error()})
			return
		}

		c.JSON(200, map[string]string {"rc": "0", "msg": "ok"})

		return
	}

	if error := model.NewSupportModel().DeleteSupportArticle(userInfo.ID, articleId); error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": error.Error()})
		return
	}

	c.JSON(200, map[string]string {"rc": "0", "msg": "ok"})
}

// 文章收藏
func (ic IndexController) CollectArticle(c *gin.Context) {
	articleIdString, isExist := c.GetPostForm("articleId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 articleId"))
	}

	articleId, error := strconv.Atoi(articleIdString)

	if error != nil {
		ic.HandleThrowException(c, errors.New("请传入 articleId"))
		return
	}

	isCollected, isExist := c.GetPostForm("isCollected")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 isCollected"))
		return
	}

	userInfo, _ := ic.GetUserInfo(c)

	if isCollected == "true" {
		if error := model.NewCollectionModel().CreateCollectArticle(userInfo.ID, articleId); error != nil {
			c.JSON(200, map[string]string {"rc": "1", "msg": error.Error()})
			return
		}

		c.JSON(200, map[string]string {"rc": "0", "msg": "ok"})

		return
	}

	if error := model.NewCollectionModel().DeleteCollectArticle(userInfo.ID, articleId); error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": error.Error()})
		return
	}

	c.JSON(200, map[string]string {"rc": "0", "msg": "ok"})
}

// 关注主播
func (ic IndexController) HandleFollowingAnchor(c *gin.Context) {
	anchorIdString, isExist := c.GetPostForm("anchorId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 anchorId"))
		return
	}

	anchorId, error := strconv.Atoi(anchorIdString)

	if error != nil {
		ic.HandleThrowException(c, error)
		return
	}

	isFollowing, isExist := c.GetPostForm("isFollowing")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 isFollowing"))
		return
	}

	userInfo, _ := ic.GetUserInfo(c)

	if isFollowing == "true" {
		if error := model.NewRelationModel().CreateUserFollowing(userInfo.ID, anchorId); error != nil {
			c.JSON(200, map[string]string {"rc": "1", "msg": error.Error()})
			return
		}

		c.JSON(200, map[string]string {"rc": "0", "msg": "ok"})

		return
	}

	if error := model.NewRelationModel().CancelUserFollowing(userInfo.ID, anchorId); error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": error.Error()})
		return
	}

	c.JSON(200, map[string]string {"rc": "0", "msg": "ok"})

}
