// package middlewares

// import (
// 	"master-gin/controllers"
// 	"master-gin/pkg/jwt"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func JWTAuthMiddleware() func(c *gin.Context) {

// 	return func(c *gin.Context) {
// 		authHeader := c.Request.Header.Get("Authorization")
// 		if authHeader == "" {
// 			controllers.ResponseError(c, http.StatusOK, controllers.CodeNeedLogin)
// 			c.Abort()
// 			return
// 		}
// 		parts := strings.SplitN(authHeader, " ", 2)
// 		if !(len(parts) == 2 && parts[0] == "Bearer") {
// 			controllers.ResponseError(c, http.StatusOK, controllers.CodeInvalidToken)
// 			c.Abort()
// 			return
// 		}
// 		mc, err := jwt.ParseToken(parts[1])
// 		if err != nil {
// 			controllers.ResponseError(c, http.StatusOK, controllers.CodeInvalidToken)
// 			c.Abort()
// 			return
// 		}
// 		c.Set(controllers.ContextUserID, mc.UserID)
// 		c.Next()
// 	}
// }
