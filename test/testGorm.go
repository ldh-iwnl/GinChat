package main

import (
	"ginchat/models"
)

func main() {

	DB.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "test"
	DB.Create(user)

	// Read
	DB.First(user, 1)
	DB.Model(user).Update("PassWord", "password")
}
