package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogoutTime    *time.Time `gorm:"column:logout_time" json:"logout_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println("v: ", v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	fmt.Println("Updating user with ID: ", user.ID)
	fmt.Println("New name: ", user.Name)
	fmt.Println("New password: ", user.PassWord)
	result := utils.DB.Model(&user).Updates(UserBasic{
		Name: user.Name, PassWord: user.PassWord})
	fmt.Println("Update result: ", result.RowsAffected)
	return result
}
