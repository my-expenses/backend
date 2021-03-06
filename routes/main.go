package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func InitializeRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	CORSMiddleware(e)
	users := e.Group("users")
	loggedIn := e.Group("/auth", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	initializeUsersRoutes(users)
	initializeTransactionsRoutes(loggedIn)
	initializeCategoriesRoutes(loggedIn)
}

func CORSMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://192.168.1.100:8080",
			"http://expenses.motawfik.com",
			"https://expenses.motawfik.com",
		},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
	}))
}
