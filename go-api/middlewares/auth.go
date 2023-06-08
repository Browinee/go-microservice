package middlewares

import (
	"go-api/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
const ContextUserID = "userId"
const ContextUserAuthorityID = "authorityId"
func JWTAuthMiddleware() func(c *gin.Context) {

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":"Please login",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// invalid token
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":"Please login",
			})
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":"Please login",
			})
			c.Abort()
			return
		}
		c.Set(ContextUserID, mc.UserId)
		c.Set(ContextUserAuthorityID, mc.AuthorityId)
		c.Next()
	}
}
