package initialize

import (
	"encoding/json"
	"fmt"
	"go-api/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func InitConfig() {
	viper.SetConfigFile("./config-dev.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig faled, err:%v\n", err)
		return
	}
	// if err := viper.Unmarshal(global.ServerConfig); err != nil {
	// 	fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	// }
	if err := viper.Unmarshal(global.NacosConfig); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	zap.S().Infof("viper naco config: %+v", global.NacosConfig)

	sc := []constant.ServerConfig{{
		IpAddr: "127.0.0.1",
		Port:   8848,
	}}
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		zap.S().Fatalf("Fail to connect naco: %+v", err.Error())
		panic(err)
	}

	content, err := configClient.GetConfig((vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group: global.NacosConfig.Group,
	}))

	if err != nil {
		zap.S().Fatalf("Fail to get config from naco: %+v", err.Error())
		panic(err)
	}
	zap.S().Infof("user service config: %+v", content)
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	zap.S().Infof("viper server config: %+v", global.ServerConfig)

 if err !=  nil {
	 zap.S().Fatalf("Fail to read naco config: %s", err.Error())
 }
	// viper.WatchConfig()

	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	fmt.Println("Config changed.....")
	// 	zap.S().Infof("viper config: %s", in.Name)
	// 	if err := viper.Unmarshal(global.NacosConfig ); err != nil {
	// 		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	// 	}
	// })

}