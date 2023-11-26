package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Summary Get all the user as a list
// @Tags get user list
// @Success 200 {string} user list data
// @Router /user/list [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// @Summary Create an user
// @Tags Create User
// @param name query string false "name"
// @param password query string false "password"
// @param rePassword query string false "rePassword"
// @Success 200 {string} user list data
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	rePassword := c.Query("rePassword")
	// get a random salt
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Salt = salt
	// validate the name
	temp := models.FindUserByName(user.Name)
	if temp.Name == user.Name {
		c.JSON(-1, gin.H{
			"message": "name is already registered",
		})
		return
	}
	if password != rePassword {
		c.JSON(-1, gin.H{
			"message": "password is not same",
		})
		return
	}
	//do encryption
	md5Pwd := utils.MakePassword(password, salt)
	// user.PassWord = password
	user.PassWord = md5Pwd
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "add successfully",
	})
}

// @Summary Delete an user
// @Tags Delete User
// @param id query string false "id"
// @Success 200 {string} msg
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))

	//find the user based on id
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "delete successfully",
	})
}

// @Summary Login as a user
// @Tags Login
// @param name query string false "name"
// @param password query string false "password"
// @Success 200 {string} user list data
// @Router /user/login [post]
func Login(c *gin.Context) {
	data := models.UserBasic{}
	// get the name and password
	name := c.Query("name")
	password := c.Query("password")
	// validate password
	user := models.FindUserByName(name)
	if user.Name != name {
		c.JSON(200, gin.H{
			"message": "user is not exist",
		})
		return
	}
	flag := utils.ValidatePassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"message": "password is not correct",
		})
		return
	}
	// login
	data = models.Login(name, utils.MakePassword(password, user.Salt))
	c.JSON(200, gin.H{
		"message": data,
	})

}

// @Summary Update an user
// @Tags Update User
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData int false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} msg
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	fmt.Println("err: ", err)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(200, gin.H{
			"message": "phone or email is not correct",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "update successfully",
		})
	}
}
