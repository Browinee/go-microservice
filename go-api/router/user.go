package router

import (
	"go-api/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func InitUserRouter(router *gin.RouterGroup){
	baseRouter := router.Group("base")
	zap.S().Info("Register base group.....")
	{
		baseRouter.GET("/captcha", api.GetCaptcha)

	}
}