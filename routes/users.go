package routes

import (
	usersController "backend/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func initializeUsersRoutes(usersRoute *echo.Group) {
	usersRoute.POST("/login", usersController.Login)
	usersRoute.POST("/refresh-token", usersController.RefreshToken,
		middleware.JWT([]byte(os.Getenv("JWT_REFRESH_TOKEN"))))
	usersRoute.POST("/register", usersController.NewUser)
	// route to check if the current JWT is valid or not
	usersRoute.GET("/status", usersController.ValidStatus,
		middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
}

