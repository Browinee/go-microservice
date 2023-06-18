package global

import (
	"grpc/config"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)