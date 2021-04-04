package routes

import (
	transactionsController "backend/controllers/transactions"
	"github.com/labstack/echo/v4"
)

func initializeTransactionsRoutes(loggedInRoute *echo.Group) {
	loggedInRoute.POST("/transactions", transactionsController.CreateTransaction)
	loggedInRoute.DELETE("/transactions", transactionsController.DeleteTransaction)
}
