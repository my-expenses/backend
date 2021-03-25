package users

import "github.com/gin-gonic/gin"

func NewUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	}
}