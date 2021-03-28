package users

import (
	customDatabaseErrors "backend/database/errors"
	usersModel "backend/models/users"
	customServicesErrors "backend/services/errors"
	usersServices "backend/services/users"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	err := usersServices.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid credentials",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func NewUser(c echo.Context) error {
	var user usersModel.User
	c.Bind(&user)

	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")
	err := usersServices.NewUser(&user, password, confirmPassword)

	if err != nil {
		if errors.Is(err, &customServicesErrors.PasswordsDontMatchError{}) {
			return c.JSON(http.StatusNotAcceptable, echo.Map{
				"message": "Passwords dont match",
			})
		}
		if errors.Is(err, &customDatabaseErrors.DuplicateEmailError{}) {
			return c.JSON(http.StatusConflict, echo.Map{
				"message": "Duplicate email",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}
	err = usersServices.NewCredentials(&usersModel.Credential{
		UserID:   user.ID,
		Password: password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})

}
