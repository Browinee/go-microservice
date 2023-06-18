package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ :=zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}