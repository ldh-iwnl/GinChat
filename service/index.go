package service

import "github.com/gin-gonic/gin"

// GetIndex

// @Tags main page
// @Accept json
// @Produce json
// @Success 200 {string} Welcome to ginchat!
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to ginchat!",
	})

}
