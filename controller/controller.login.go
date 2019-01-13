package controller

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/base64"
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

	c.JSON(200, map[string]string {"captchaId": captchaId, "data": "data:image/png;base64," + captchaBase64String})
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

}

// 忘记密码
func (lc LoginController) FindPassword(c *gin.Context) {

}
