package controller

import (
	"github.com/gin-gonic/gin"
	"yuedu/model"
)

type UserController struct {
	Controller
}

func (uc UserController) HandleFollows(c *gin.Context) {
	followers := model.NewRelationModel().GetUserFollowers(256)

	c.JSON(200, followers)
}

func (uc UserController) HanleFollowings(c *gin.Context) {
	followers := model.NewRelationModel().GetUserFollowings(256)

	c.JSON(200, followers)
}
