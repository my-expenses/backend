package routes

import (
	transactionsController "backend/controllers/transactions"
	"github.com/labstack/echo/v4"
)

func initializeTransactionsRoutes(loggedInRoute *echo.Group) {
	loggedInRoute.GET("/transactions", transactionsController.GetTransactions)
	loggedInRoute.POST("/transactions", transactionsController.CreateTransaction)
	loggedInRoute.DELETE("/transactions/:transactionID", transactionsController.DeleteTransaction)
}
