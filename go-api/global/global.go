package global

import (
	"go-api/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
)

// NOTE:global variables
var (
	 RedisClient *redis.Client
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans ut.Translator
)