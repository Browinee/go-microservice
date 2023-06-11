package api

import (
	"fmt"
	"go-api/forms"
	"go-api/global"
	"go-api/utils"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)


func GenerateSmsCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i :=0; i<width;i++ {
			fmt.Fprintf(&sb,"%d", numeric[rand.Intn(r)])
	}
		return sb.String()
}
func SendSms(ctx *gin.Context){
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBindJSON(&sendSmsForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
    zap.S().Errorf("valid %v", errs.Translate(global.Trans))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.RemoveTopStruct(errs.Translate(global.Trans)),
		})
		return
	}
	// 這裡需要註冊send sms服務，暫略

	// sms發送後，要從response中取的code並加入redis中
  smsCode := GenerateSmsCode(6)
	zap.S().Infof("smsCode: %s", smsCode)
	mobile:="12312312312"
	global.RedisClient.Set(mobile, smsCode, 300 * time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":"Send sms code successfully",

	})
}

