package routes

import (
	transactionsController "backend/controllers/transactions"
	"github.com/labstack/echo/v4"
)

func initializeTransactionsRoutes(loggedInRoute *echo.Group) {
	loggedInRoute.GET("/transactions", transactionsController.GetTransactions)
	loggedInRoute.GET("/grouped-transactions", transactionsController.GetGroupedTransactions)
	loggedInRoute.POST("/transactions", transactionsController.CreateTransaction)
	loggedInRoute.PUT("/transactions/:transactionID", transactionsController.UpdateTransaction)
	loggedInRoute.DELETE("/transactions/:transactionID", transactionsController.DeleteTransaction)
}
