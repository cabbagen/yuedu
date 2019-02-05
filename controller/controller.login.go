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
)

type LoginController struct {
	Controller
}

// 用户登录
func (lc LoginController) Login(c *gin.Context) {

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
		c.JSON(200, map[string]string {"rc": "1", "msg": "生成验证码错误"})
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
func (lc LoginController) ValidateCaptcha(c *gin.Context) {
	captchaId, hasCaptchaId := c.GetQuery("captchaId")

	captchaDigits, hasCaptchaDigits := c.GetQuery("captchaDigits")

	if !hasCaptchaId {
		c.JSON(200, map[string]string {"rc": "1", "msg": "captchaId 不存在"})
		return
	}

	if !hasCaptchaDigits {
		c.JSON(200, map[string]string {"rc": "1", "msg": "hasCaptchaDigits 不存在"})
		return
	}

	var isValidated bool = captcha.VerifyString(captchaId, captchaDigits)

	if !isValidated {
		c.JSON(200, map[string]string {"rc": "1", "data": "验证失败"})
		return
	}

	c.JSON(200, map[string]string {"rc": "0", "data": "验证通过"})
}

// 用户注册
func (lc LoginController) Register(c *gin.Context) {

	username, hasUsername := c.GetPostForm("username")

	if !hasUsername {
		c.JSON(200, map[string]string {"rc": "1", "msg": "参数错误"})
		return
	}

	password, hasPassword := c.GetPostForm("password")

	if !hasPassword {
		c.JSON(200, map[string]string {"rc": "1", "msg": "参数错误"})
		return
	}

	email, hasEmail := c.GetPostForm("email")

	if !hasEmail {
		c.JSON(200, map[string]string {"rc": "1", "msg": "参数错误"})
		return
	}

	var user schema.User = schema.User {
		UserName: username,
		PassWord: utils.MakeMD5(password),
		Email: email,
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	var isSuccess bool = model.NewUserModel().CreateUserInfo(user)

	if !isSuccess {
		c.JSON(200, map[string]string {"rc": "1", "msg": "当前用户已存在"})
		return
	}

	c.JSON(200, map[string]string {"rc": "0", "msg": "用户注册成功"})
}

// 忘记密码
func (lc LoginController) FindPassword(c *gin.Context) {

}
