package controller

import (
	"github.com/gin-gonic/gin"
	"yuedu/model"
	"strconv"
)

type SearchController struct {
	Controller
}

func (sc SearchController) RenderSearch(c *gin.Context) {

	keyword := c.DefaultQuery("keyword", "")

	searchData, error := sc.GetSearchData(keyword)

	if error != nil {
		sc.HandleThrowException(c, error)
	}

	if userInfo, isExist := sc.GetUserInfo(c); isExist {
		searchData["userInfo"] = userInfo;
	}

	c.HTML(200, "wsearch.html", searchData)
}

func (sc SearchController) GetSearchData(keyword string) (map[string]interface{}, error) {
	var searchData map[string]interface{} = make(map[string]interface{})

	articleModel := model.NewArticleModel()

	if channels, error := model.NewChannelModel().GetAllChannels(); error != nil {
		return searchData, error
	} else {
		searchData["channels"] = channels
	}

	if articles, error := articleModel.GetArticlesByKeyword(keyword, 0, 10); error != nil {
		return searchData, error
	} else {
		searchData["articles"] = articles
	}

	if topArticles, error := articleModel.GetCommonLastArticles(10); error != nil {
		return searchData, error
	} else {
		searchData["topArticles"] = topArticles
	}

	if count, error := articleModel.GetArticleCountByKeyword(keyword); error != nil {
		return searchData, error
	} else {
		searchData["count"] = count
	}

	return searchData, nil
}

func (sc SearchController) GetSearchArticles(c *gin.Context) {

	keyword := c.DefaultQuery("keyword", "")


	page, error := strconv.Atoi(c.DefaultQuery("page", "0"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	size, error := strconv.Atoi(c.DefaultQuery("size", "10"))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请求参数错误"})
		return
	}

	articleModel := model.NewArticleModel()

	articles, error := articleModel.GetArticlesByKeyword(keyword, page, size)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "2", "msg": error.Error()})
		return
	}

	count, error := articleModel.GetArticleCountByKeyword(keyword)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "2", "msg": error.Error()})
		return
	}

	c.JSON(200, map[string]interface{} {"articles": articles, "count": count, "rc": "0"})
}