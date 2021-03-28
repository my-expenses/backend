package routes

import (
	usersRoutes "backend/routes/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitializeRoutes(e *echo.Echo) {
	CORSMiddleware(e)
	users := e.Group("users")
	usersRoutes.InitializeUsersRoutes(users)
}

func CORSMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://192.168.1.100:8080",
		},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
	}))
}
