package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb3&parseTime=True&loc=Local"), &gorm.Config{})
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "test"
	db.Create(user)

	// Read
	db.First(user, 1)
	db.Model(user).Update("PassWord", "password")
}
