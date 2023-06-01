package initialize

import (
	"go-api/api"

	"go.uber.org/zap"
)



func InitValidator() {
	if err := api.InitTrans("en"); err != nil {
		zap.S().Errorf("init validator trans failed %v\n", err)
		return
	}
}