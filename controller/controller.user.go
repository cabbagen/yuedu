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
		uc.HandleThrowException(c, error)
		return
	}

	userInfo, error := model.NewUserModel().GetUserInfo(userId)

	if error != nil {
		uc.HandleThrowException(c, error)
		return
	}

	c.HTML(200, "wuser.html", userInfo)
}

//func (uc UserController) GetUserData(userId int) (map[string]interface{}, error) {

	//var userData map[string]interface{} = make(map[string]interface{})

	//userInfo := model.NewUserModel().GetUserInfo(userId)
	//
	//collectedArticles, error := model.NewCollectionModel().GetUserCollectedArticles(userId)
	//
	//
	//
	//if error != nil {
	//	return map[string]interface{} {""}
	//}
//}

func (uc UserController) HandleFollows(c *gin.Context) {
	//followers := model.NewRelationModel().GetUserFollowers(256)
	//
	//c.JSON(200, followers)
}

func (uc UserController) HandleFollowings(c *gin.Context) {
	//followers := model.NewRelationModel().GetUserFollowings(256)
	//
	//c.JSON(200, followers)
}
