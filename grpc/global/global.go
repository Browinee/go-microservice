package global

import (
	"grpc/config"
	"grpc/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
	ServerConfig *config.ServerConfig
)
func init(){

}