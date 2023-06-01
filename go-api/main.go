// NOTE: gin
package main

import (
	"fmt"
	"go-api/global"
	"go-api/initialize"

	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
  initialize.InitConfig()
  initialize.InitValidator()

	Router := initialize.Routers()
	zap.S().Infof("Start server port %d...", global.ServerConfig.Port)
	if err :=	Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("Fail to start server", err.Error() )
	}
}
