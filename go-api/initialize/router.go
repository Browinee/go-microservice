package initialize

import (
	"go-api/router"

	"github.com/gin-gonic/gin"
)




func Routers() *gin.Engine {
	defaultRouter := gin.Default()
	apiGroup :=defaultRouter.Group("/v1")
	router.InitUserRouter(apiGroup)

	return defaultRouter
}