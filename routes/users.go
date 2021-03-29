package routes

import (
	usersController "backend/controllers/users"
	"github.com/labstack/echo/v4"
)

func initializeUsersRoutes(usersRoute *echo.Group) {
	usersRoute.POST("/login", usersController.Login)
	usersRoute.POST("/register", usersController.NewUser)
}

