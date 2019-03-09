package controller

import (
	"github.com/gin-gonic/gin"
	"yuedu/model"
	"strconv"
	"log"
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

	userData, error := uc.GetUserData(userId)

	if error != nil {
		uc.HandleThrowException(c, error)
		return
	}

	c.HTML(200, "wuser.html", userData)
}

func (uc UserController) GetUserData(userId int) (map[string]interface{}, error) {

	var userData map[string]interface{} = make(map[string]interface{})

	// 用户信息
	if userInfo, error := model.NewUserModel().GetUserInfo(userId); error != nil {
		return userData, error
	} else {
		userData["userInfo"] = userInfo
	}

	// 用户收藏文章
	if articles, error := model.NewCollectionModel().GetUserCollectedArticles(userId); error != nil {
		return userData, error
	} else {
		userData["articles"] = articles
	}

	// 用户粉丝
	if followers, error := model.NewRelationModel().GetUserFollowers(userId); error != nil {
		return userData, nil
	} else {
		userData["followers"] = followers
	}

	// 用户关注
	if followings, error := model.NewRelationModel().GetUserFollowings(userId); error != nil {
		return userData, nil
	} else {
		userData["followings"] = followings
	}

	return userData, nil
}

func (uc UserController) FetchUserData(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))

	userData, error := uc.GetUserData(userId)

	if error != nil {
		log.Println(error)
	}

	c.JSON(200, userData)
}

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
