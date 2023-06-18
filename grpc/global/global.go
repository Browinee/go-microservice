package global

import (
	"grpc/config"

	"gorm.io/gorm"
)

// import (
// 	"grpc/config"
// 	"grpc/model"
// 	"log"
// 	"os"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

var (
	DB *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
// func init(){
// 	dsn := "root:root1234@tcp(localhost:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags),
// 		logger.Config{
// 			SlowThreshold: time.Second,
// 			 LogLevel: logger.Info,
// 			 Colorful: true,
// 		},
// 	)
// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger:newLogger,
// 	})
// if err != nil {
// 	panic(err)
// }
// _ = DB.AutoMigrate(&model.User{})
// }