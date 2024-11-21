package api

import (
	"context"
	"ginWebApi/models"
	pbCaptcha "ginWebApi/proto/captcha"
	"log"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct{}

func (con CaptchaController) MakeCaptcha(c *gin.Context) {
	// Create client
	captchaClient := pbCaptcha.NewCaptchaService("captcha", models.CaptchaClient)
	// Call service
	res, err := captchaClient.MakeCaptcha(context.Background(), &pbCaptcha.MakeCaptchaRequest{
		Height: 100,
		Width:  280,
		Length: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"captchaId": res.Id,
		"B64s":      res.B64S,
	})
}

func (con CaptchaController) VerifyCaptcha(c *gin.Context) {
	verifyId := c.PostForm("verifyId")
	verifyValue := c.PostForm("verifyValue")

	captchaClient := pbCaptcha.NewCaptchaService("captcha", models.CaptchaClient)

	res, _ := captchaClient.VerifyCaptcha(context.Background(), &pbCaptcha.VerifyCaptchaRequest{
		Id:          verifyId,
		VerifyValue: verifyValue,
	})
	if res.VerifyResult == true {
		c.JSON(200, gin.H{
			"message": "验证码验证成功",
			"success": true,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "验证码验证失败",
			"success": false,
		})
	}
}
