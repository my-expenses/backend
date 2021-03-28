package users

import (
	usersController "backend/controllers/users"
	"github.com/labstack/echo/v4"
)

func InitializeUsersRoutes(usersRoute *echo.Group) {
	usersRoute.POST("/login", usersController.Login)
	usersRoute.POST("/register", usersController.NewUser)
}

