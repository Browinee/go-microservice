package router

import (
	"go-api/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func InitUserRouter(router *gin.RouterGroup){
	userRouter := router.Group("user")
	zap.S().Info("Register user group.....")
	{
		userRouter.GET("/", api.GetUserList)

	}
}