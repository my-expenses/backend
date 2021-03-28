package users

import (
	customDatabaseErrors "backend/database/errors"
	usersModel "backend/models/users"
	customServicesErrors "backend/services/errors"
	usersServices "backend/services/users"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	err := usersServices.Login(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func NewUser(c *gin.Context) {
	var user usersModel.User
	c.Bind(&user)

	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	err := usersServices.NewUser(&user, password, confirmPassword)

	if err != nil {
		if errors.Is(err, &customServicesErrors.PasswordsDontMatchError{}) {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": "Passwords dont match",
			})
			return
		}
		if errors.Is(err, &customDatabaseErrors.DuplicateEmailError{}) {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Duplicate email",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	err = usersServices.NewCredentials(&usersModel.Credential{
		UserID:   user.ID,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}