// NOTE: gin
package main

import (
	"fmt"
	"go-api/global"
	"go-api/initialize"
	myValidator "go-api/validators"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
  initialize.InitConfig()
  initialize.InitValidator()
	initialize.InitSrcConn()
  redisErr := initialize.InitRedis()
	if redisErr!= nil {
		zap.S().Errorf("init redis failed, err %v\n", redisErr)
	}
	defer initialize.CloseRedis()


	Router := initialize.Routers()
  // NOTE: 非local的話，自動獲取
	// port, err := utils.GetFreePort()
	// if err ==nil {
	// 	global.ServerConfig.Port = port
	// }
  if v , ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", myValidator.ValidateMobile)
	}


	zap.S().Infof("Start server port %d...", global.ServerConfig.Port)
	if err :=	Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("Fail to start server", err.Error() )
	}
}
