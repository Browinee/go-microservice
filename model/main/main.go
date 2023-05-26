package main

import (
	"crypto/md5"
	"encoding/hex"
	"go-microservice/model"
	"log"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)





func generatePassword(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charactersLength := len(characters)
	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charactersLength)
		password[i] = characters[randomIndex]
	}

	return string(password)
}
func main(){
// 	dsn := "root:root1234@tcp(localhost:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags),
// 		logger.Config{
// 			SlowThreshold: time.Second,
// 			 LogLevel: logger.Info,
// 			 Colorful: true,
// 		},
// 	)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger:newLogger,
// 	})
// if err != nil {
// 	panic(err)
// }
// _ = db.AutoMigrate(&model.User{})
}