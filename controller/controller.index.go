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

  var channels []schema.Channel
  model.NewChannelModel().FindAll(&channels)

  var article schema.Article
  model.NewArticleModel().GetLastArticleByChannel(&article, 1)

  var userInfo model.FullUserInfo
  model.NewUserModel().GetFullUserInfo(article.Anchor, &userInfo)

  var indexData map[string]interface{} = map[string]interface{} {
    "channels": channels,
    "article": article,
    "userInfo": userInfo,
  }

  return indexData
  
}
