package initialize

import (
	"fmt"
	"go-api/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (

	Nil    = redis.Nil
)

func InitRedis() (err error) {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", global.ServerConfig.Redis.HOST, global.ServerConfig.Redis.Port),
		Password:    global.ServerConfig.Redis.Password,
		DB:           global.ServerConfig.Redis.DB,
		PoolSize:     global.ServerConfig.Redis.PoolSize,
		MinIdleConns:  global.ServerConfig.Redis.MinIdleConns,
	})
	_, err = 	global.RedisClient.Ping().Result()
	if err != nil {
		zap.L().Error("connect redis failed, err %v/n", zap.Error(err))
		return
	}
	zap.S().Info("Init redis successfully.")
	return err
}

func CloseRedis() {
	_ = 	global.RedisClient.Close()
}
