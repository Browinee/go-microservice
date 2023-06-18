package initialize

import (
	"fmt"
	"grpc/global"
	"grpc/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() {
	mysqlInfo := global.ServerConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local", mysqlInfo.User, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Name)
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
_ = DB.AutoMigrate(&model.User{})
}