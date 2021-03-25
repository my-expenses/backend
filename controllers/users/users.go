package users

import (
	customErrors "backend/database/errors"
	usersModel "backend/models/users"
	usersServices "backend/services/users"
	"errors"
	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	var user usersModel.User
	c.Bind(&user)

	password := c.PostForm("password")
	err := usersServices.NewUser(&user, password)

	if err != nil {
		if errors.Is(err, &customErrors.DuplicateEmailError{}) {
			c.JSON(409, gin.H{
				"message": "Duplicate email",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})

}