package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func IsAdminAuth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authorityId, _ := ctx.Get(ContextUserAuthorityID)
		zap.S().Infof("authorityId: %v", authorityId)
		if authorityId.(int32) == 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg":"forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}