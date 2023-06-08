package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)


var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context){
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		zap.S().Errorf("Fail to generate captcha", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":"Fail to generate captcha",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath": b64s,
	})
}
