package users

import (
	customDatabaseErrors "backend/database/errors"
	authMiddleware "backend/middlewares/auth"
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
	userID, err := usersServices.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid credentials",
		})
	}
	token, err := authMiddleware.GenerateJWT(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error generating the token",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"message": "success",
	})
}

func NewUser(c echo.Context) error {
	var user usersModel.User
	c.Bind(&user)
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var credentials usersModel.Credential
	c.Bind(&credentials)
	if err := c.Validate(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
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
