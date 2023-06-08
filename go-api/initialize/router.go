package initialize

import (
	"go-api/middlewares"
	"go-api/router"

	"github.com/gin-gonic/gin"
)




func Routers() *gin.Engine {
	defaultRouter := gin.Default()
	defaultRouter.Use(middlewares.Cors())
	apiGroup :=defaultRouter.Group("/v1")
	router.InitUserRouter(apiGroup)
	router.InitBaseRouter(apiGroup)

	return defaultRouter
}