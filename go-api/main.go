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
	Router := initialize.Routers()
	defer initialize.CloseRedis()

  if v , ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", myValidator.ValidateMobile)
	}


	zap.S().Infof("Start server port %d...", global.ServerConfig.Port)
	if err :=	Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("Fail to start server", err.Error() )
	}
}
