package users

import (
	customDatabaseErrors "backend/database/errors"
	authMiddleware "backend/middlewares/auth"
	usersModel "backend/models/users"
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
	accessToken, accessErr := authMiddleware.GenerateJWT(userID, true)
	refreshToken, refreshErr := authMiddleware.GenerateJWT(userID, false)
	if accessErr != nil || refreshErr != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error generating the accessToken",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"accessToken":   accessToken,
		"refreshToken": refreshToken,
		"message": "success",
	})
}

func RefreshToken(c echo.Context) error {
	userID := authMiddleware.FetchLoggedInUserID(&c)
	accessToken, accessErr := authMiddleware.GenerateJWT(userID, true)
	refreshToken, refreshErr := authMiddleware.GenerateJWT(userID, false)
	if accessErr != nil || refreshErr != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error generating the accessToken",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"accessToken":   accessToken,
		"refreshToken": refreshToken,
		"message": "success",
	})
}

func NewUser(c echo.Context) error {
	var user usersModel.User
	c.Bind(&user)
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	var credentials usersModel.Credential
	c.Bind(&credentials)
	if err := c.Validate(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err := usersServices.NewUser(&user)
	if err != nil {
		if errors.Is(err, &customDatabaseErrors.DuplicateEmailError{}) {
			return c.JSON(http.StatusConflict, echo.Map{
				"message": "Duplicate email",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}
	credentials.User = user
	err = usersServices.NewCredentials(&credentials)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})

}
