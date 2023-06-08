package router

import (
	"go-api/api"
	"go-api/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func InitBaseRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	zap.S().Info("Register user group.....")
	{
		userRouter.POST("/login",  api.PassWordLogin)
		userRouter.GET("/list", middlewares.JWTAuthMiddleware(),middlewares.IsAdminAuth(), api.GetUserList)

	}
}