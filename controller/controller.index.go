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

		// 用户收藏点赞信息
		if extraInfo, error := ic.GetUserExtraInfo(userInfo.ID, 354); error == nil {
			indexData["extraInfo"] = extraInfo
		}

		// 用户评论信息
		if comments, error := ic.GetArticleComments(userInfo.ID, 354); error == nil {
			indexData["comments"] = comments
		}
	} else {
		indexData["comments"] = []interface{}{}
	}

	c.HTML(200, "windex.html", indexData)
}

func (ic IndexController) GetIndexData(articleId int) (map[string]interface{}, error) {

	var indexData map[string]interface{} = make(map[string]interface{})

	channelModel, articleModel := model.NewChannelModel(), model.NewArticleModel()

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

func (ic IndexController) GetArticleComments(userId, articleId int) ([]model.CommentInfo, error) {
	comments, error := model.NewCommentlModel().GetArticleComments(userId, articleId)

	if error != nil {
		return comments, error
	}

	return comments, nil
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

		// 用户收藏点赞信息
		if extraInfo, error := ic.GetUserExtraInfo(userInfo.ID, articleId); error == nil {
			detailData["extraInfo"] = extraInfo
		}

		// 用户评论信息
		if comments, error := ic.GetArticleComments(userInfo.ID, articleId); error == nil {
			detailData["comments"] = comments
		}
	} else {
		detailData["comments"] = []interface{}{}
	}

	c.HTML(200, "windex.html", detailData)
}


// 文章点赞
func (ic IndexController) SupportArticle(c *gin.Context) {
	articleIdString, isExist := c.GetPostForm("articleId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 articleId"))
	}

	articleId, _ := strconv.Atoi(articleIdString)

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

	articleId, _ := strconv.Atoi(articleIdString)

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

	anchorId, _ := strconv.Atoi(anchorIdString)

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

// 处理文章评论
func (ic IndexController) HandleArticleComment(c *gin.Context) {
	articleIdString, isExist := c.GetPostForm("articleId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 articleId"))
		return
	}

	articleId, _ := strconv.Atoi(articleIdString)

	comment, isExist := c.GetPostForm("comment")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 comment"))
		return
	}

	userInfo, _ := ic.GetUserInfo(c)

	error := model.NewCommentlModel().CreateArticleComment(userInfo.ID, articleId, comment)

	if error != nil {
		ic.HandleThrowException(c, error)
		return
	}

	c.JSON(200, map[string]string { "rc": "0", "msg": "ok" })
}

// 处理评论评论
func (ic IndexController) HandleCommentComment(c *gin.Context) {
	commentIdString, isExist := c.GetPostForm("commentId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 commentId"))
		return
	}

	commentId, _ := strconv.Atoi(commentIdString)

	comment, isExist := c.GetPostForm("comment")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 comment"))
		return
	}

	userInfo, _ := ic.GetUserInfo(c)

	error := model.NewCommentlModel().CreateCommentComment(userInfo.ID, commentId, comment)

	if error != nil {
		ic.HandleThrowException(c, error)
		return
	}

	c.JSON(200, map[string]string { "rc": "0", "msg": "ok" })
}

// 删除评论
func (ic IndexController) HandleDeleteComment(c *gin.Context) {
	commentIdString, isExist := c.GetPostForm("commentId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 commentId"))
		return
	}

	commentId, _ := strconv.Atoi(commentIdString)

	error := model.NewCommentlModel().DeleteComment(commentId)

	if error != nil {
		ic.HandleThrowException(c, error)
		return
	}

	c.JSON(200, map[string]string { "rc": "0", "msg": "ok" })
}

// 评论点赞
func (ic IndexController) HandleCommentSupport(c *gin.Context) {
	commentIdString, isExist := c.GetPostForm("commentId")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 commentId"))
		return
	}

	commentId, _ := strconv.Atoi(commentIdString)

	isSupport, isExist := c.GetPostForm("isSupport")

	if !isExist {
		ic.HandleThrowException(c, errors.New("请传入 isSupport"))
		return
	}

	userInfo, _ := ic.GetUserInfo(c)

	if isSupport == "true" {
		if error := model.NewSupportModel().CreateSupportComment(userInfo.ID, commentId); error != nil {
			ic.HandleThrowException(c, error)
			return
		}
		c.JSON(200, map[string]string { "rc": "0", "msg": "ok" })
		return
	}

	if error := model.NewSupportModel().DeleteSupportComment(userInfo.ID, commentId); error != nil {
		ic.HandleThrowException(c, error)
		return
	}

	c.JSON(200, map[string]string { "rc": "0", "msg": "ok" })
}

