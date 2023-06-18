package initialize

import (
	"fmt"
	"grpc/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)
func InitConfig() {
	viper.SetConfigFile("./config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig faled, err:%v\n", err)
		return
	}
	if err := viper.Unmarshal(global.ServerConfig); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	zap.S().Infof("viper config: %v", global.ServerConfig)
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config changed.....")
		zap.S().Infof("viper config: %s", in.Name)

		if err := viper.Unmarshal(global.ServerConfig); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}