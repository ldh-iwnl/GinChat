package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all the user as a list
// @Tags get user list
// @Accept json
// @Produce json
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
	if password != rePassword {
		c.JSON(-1, gin.H{
			"message": "password is not same",
		})
		return
	}
	user.PassWord = password
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

// @Summary Update an user
// @Tags Update User
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData int false "password"
// @Success 200 {string} msg
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "update successfully",
	})
}
