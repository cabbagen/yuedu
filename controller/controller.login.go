package controller

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/base64"
	"yuedu/schema"
	"yuedu/utils"
	"yuedu/model"
	"time"
	"yuedu/middleware"
	"encoding/json"
)

type LoginController struct {
	Controller
}

// 用户登录
func (lc LoginController) Login(c *gin.Context) {
	username, hasUsername := c.GetPostForm("username")

	if !hasUsername {
		c.JSON(200, map[string]string{"rc": "1", "msg": "请输入用户名"})
		return
	}

	password, hasPassword := c.GetPostForm("password")

	if !hasPassword {
		c.JSON(200, map[string]string{"rc": "1", "msg": "请输入密码"})
		return
	}

	captchaId, hasCaptchaId := c.GetPostForm("captchaId")

	if !hasCaptchaId {
		c.JSON(200, map[string]string {"rc": "1", "msg": "captchaId 不存在"})
		return
	}

	captchaDigits, hasCaptchaDigits := c.GetPostForm("captchaDigits")

	if !hasCaptchaDigits {
		c.JSON(200, map[string]string {"rc": "1", "msg": "hasCaptchaDigits 不存在"})
		return
	}

	if isValid := lc.ValidateCaptcha(captchaId, captchaDigits); !isValid {
		c.JSON(200, map[string]string {"rc": "2", "msg": "验证码已过期"})
		return
	}

	userModel := model.NewUserModel()

	isExist, error := userModel.ValidateUserInfo(username, password)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "5", "msg": "数据库操作失败: " + error.Error()})
		return
	}

	if !isExist {
		c.JSON(200, map[string]string {"rc": "3", "msg": "用户名密码错误"})
		return
	}

	userInfo, error := userModel.GetUserInfoByName(username)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "5", "msg": "数据库操作失败: " + error.Error()})
		return
	}

	userInfoJson, error := json.Marshal(userInfo)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "4", "msg": "token 生成错误: " + error.Error()})
		return
	}

	token, error := middleware.NewTokenMiddleware().SignToken(string(userInfoJson))

	if error != nil {
		c.JSON(200, map[string]string {"rc": "4", "msg": "token 生成错误: " + error.Error()})
		return
	}

	userInfoString, maxAge := utils.AESECBEncode(utils.AESKEY, string(userInfoJson)), int((time.Hour * 24).Seconds())

	c.SetCookie("userInfo", userInfoString, maxAge, "/", "localhost", false, true)

	c.JSON(200, map[string]string {"rc": "0", "data": token})
}

// 生成验证码
func (lc LoginController) GetCaptcha(c *gin.Context) {
	captchaId, hasCaptchaId := c.GetQuery("captchaId")

	if hasCaptchaId {
		captcha.Reload(captchaId)
	} else {
		captchaId = captcha.New()
	}

	var imageBuffer *bytes.Buffer = bytes.NewBuffer([]byte{})

	if error := captcha.WriteImage(imageBuffer, captchaId, captcha.StdWidth, captcha.StdHeight); error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "生成验证码错误, 请刷新页面重试"})
		return
	}

	var captchaBase64String = base64.StdEncoding.EncodeToString(imageBuffer.Bytes())

	c.JSON(200, map[string]interface{} {
		"rc": "0",
		"data": map[string]string {
			"captchaId": captchaId,
			"img": "data:image/png;base64," + captchaBase64String,
		},
	})
}

// 校验验证码
func (lc LoginController) ValidateCaptcha(captchaId, captchaDigits string) bool {
	return captcha.VerifyString(captchaId, captchaDigits)
}

// 用户注册
func (lc LoginController) HandleRegister(c *gin.Context) {

	username, hasUsername := c.GetPostForm("username")

	if !hasUsername {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请输入用户名"})
		return
	}

	password, hasPassword := c.GetPostForm("password")

	if !hasPassword {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请输入用户密码"})
		return
	}

	email, hasEmail := c.GetPostForm("email")

	if !hasEmail {
		c.JSON(200, map[string]string {"rc": "1", "msg": "请输入用户邮箱"})
		return
	}

	user := schema.User {
		Username: username,
		Password: utils.MakeMD5(password),
		Email: email,
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	error := model.NewUserModel().CreateUserInfo(user)

	if error != nil {
		c.JSON(200, map[string]string {"rc": "1", "msg": "当前用户已存在"})
		return
	}

	c.JSON(200, map[string]string {"rc": "0", "msg": "用户注册成功"})
}

// 忘记密码
func (lc LoginController) FindPassword(c *gin.Context) {

}
