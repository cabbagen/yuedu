package controller

import (
  "../model"
  "../schema"
  "github.com/gin-gonic/gin"
)

type IndexController struct {
  Controller
}

func (ic IndexController) HandleIndex(c *gin.Context) {
  var users []schema.User;
  var userModel model.UserModel = model.NewUserModel();

  userModel.Query(&users)

  c.JSON(200, users)

  // c.HTML(200, "index.tmpl", map[string]string {"title": "hello golang"})
}
