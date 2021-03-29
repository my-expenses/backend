package routes

import (
	categoriesController "backend/controllers/transactions"
	"github.com/labstack/echo/v4"
)

func initializeTransactionsRoutes(loggedInRoute *echo.Group) {
	loggedInRoute.GET("/categories", categoriesController.GetCategories)
	loggedInRoute.POST("/categories", categoriesController.CreateCategory)
	loggedInRoute.DELETE("/categories/:categoryID", categoriesController.DeleteCategory)
}
