package controller

import (
	"github.com/gin-gonic/gin"
	"yuedu/model"
	"strconv"
)

type UserController struct {
	Controller
}

func (uc UserController) RenderUser(c *gin.Context) {

	userId, error := strconv.Atoi(c.Param("userId"))

	if error != nil {
		uc.HanleThrowException(c, error)
		return
	}

	userInfo := model.NewUserModel().GetUserInfo(userId)

	c.HTML(200, "wuser.html", userInfo)
}

func (uc UserController) HandleFollows(c *gin.Context) {
	followers := model.NewRelationModel().GetUserFollowers(256)

	c.JSON(200, followers)
}

func (uc UserController) HandleFollowings(c *gin.Context) {
	followers := model.NewRelationModel().GetUserFollowings(256)

	c.JSON(200, followers)
}
