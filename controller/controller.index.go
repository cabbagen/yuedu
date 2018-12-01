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
  var indexData map[string]interface{} = make(map[string]interface{})

  var channels []schema.Channel
  model.NewChannelModel().FindAll(&channels)

  var article schema.Article
  model.NewArticleModel().GetLastArticleByChannel(&article, 1)

  var userInfo model.FullUserInfo
  model.NewUserModel().GetFullUserInfo("168", &userInfo)

  indexData["channels"] = channels
  indexData["article"] = article
  indexData["userInfo"] = userInfo

  c.PureJSON(200, indexData)

  // c.HTML(200, "windex.html", map[string]string {"title": "hello golang"})
}
