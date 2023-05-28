// NOTE: gin
package main

import (
	"fmt"
	"go-microservice/initialize"

	"go.uber.org/zap"
)

func main() {
	port := 8085
	initialize.InitLogger()

	Router := initialize.Routers()

	zap.S().Infof("Start server port %d...", port)
	if err :=	Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("Fail to start server", err.Error() )
	}
}
