package initialize

import (
	"fmt"
	"grpc/global"
	"grpc/model"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() {
	mysqlInfo := global.ServerConfig.Mysql
	zap.S().Infof("mysql %#v", mysqlInfo)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DB)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			 LogLevel: logger.Info,
			 Colorful: true,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:newLogger,
	})
if err != nil {
	panic(err)
}
_ = global.DB.AutoMigrate(&model.User{})
}