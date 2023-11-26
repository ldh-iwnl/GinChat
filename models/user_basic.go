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
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIP      string
	ClientPort    string
	Salt          string
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

func Login(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) *gorm.DB {
	return utils.DB.Where("phone = ?", phone).First(&UserBasic{})
}

func FindUserByEmail(email string) *gorm.DB {
	return utils.DB.Where("email = ?", email).First(&UserBasic{})
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
	fmt.Println("New phone: ", user.Phone)
	fmt.Println("New email: ", user.Email)
	result := utils.DB.Model(&user).Updates(UserBasic{
		Name: user.Name, PassWord: user.PassWord, Phone: user.Phone,
		Email: user.Email})
	fmt.Println("Update result: ", result.RowsAffected)
	return result
}
