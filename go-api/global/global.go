package global

import (
	"go-api/config"
	"go-api/proto"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
)

// NOTE:global variables
var (
	 RedisClient *redis.Client
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans ut.Translator
	UserSrvClient proto.UserClient
	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)