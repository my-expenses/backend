package routes

import (
	transactionsController "backend/controllers/transactions"
	"github.com/labstack/echo/v4"
)

func initializeTransactionsRoutes(loggedInRoute *echo.Group) {
	loggedInRoute.POST("/categories", transactionsController.CreateTransaction)
	loggedInRoute.DELETE("/categories", transactionsController.DeleteTransaction)
}
