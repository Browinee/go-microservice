package main

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
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


	// Using custom options
	options := &password.Options{10, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)

	fmt.Println(check) // true
  // NOTE: gen user
	// for i :=0; i<10;i++ {
	// 	user := model.User{
	// 		NickName: fmt.Sprintf("Test-%d", i),
	// 		Mobile: fmt.Sprintf("18788822%d", i),
	// 		Password: newPassword,
	// 	}
	// 	global.DB.Save(&user)
	// }
}