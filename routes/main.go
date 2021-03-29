package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func InitializeRoutes(e *echo.Echo) {
	CORSMiddleware(e)
	users := e.Group("users")
	loggedIn := e.Group("/auth", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	initializeUsersRoutes(users)
	initializeTransactionsRoutes(loggedIn)
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
