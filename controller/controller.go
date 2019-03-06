package controller

import (
	"yuedu/schema"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"yuedu/utils"
)

type Controller struct {
}

func (cl Controller) GetUserInfo(c *gin.Context) (schema.User, bool) {
	userInfo := schema.User{}

	userInfoString, error := c.Cookie("userInfo")

	if error != nil {
		return userInfo, false
	}

	userInfoJson := utils.AESECBDecode(utils.AESKEY, userInfoString)


	if error := json.Unmarshal([]byte(userInfoJson), &userInfo); error != nil {
		return userInfo, false
	}

	return userInfo, true
}

func (cl Controller) HanleThrowException(c *gin.Context, err error) {
	c.JSON(200, map[string]string {"rc": "1", "msg": err.Error()})
}
