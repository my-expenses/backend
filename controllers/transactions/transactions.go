package transactions

import (
	"backend/controllers/pagination"
	customDatabaseErrors "backend/database/errors"
	authController "backend/middlewares/auth"
	transactionsModel "backend/models/transactions"
	transactionsServices "backend/services/transactions"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func CreateTransaction(c echo.Context) error {
	var transaction transactionsModel.Transaction
	c.Bind(&transaction)
	err := c.Validate(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	transaction.UserID = authController.FetchLoggedInUserID(&c)
	err = transactionsServices.CreateTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"transaction": transaction,
	})
}

func UpdateTransaction(c echo.Context) error {
	var transaction transactionsModel.Transaction
	c.Bind(&transaction)
	transactionID := c.Param("transactionID")
	err := c.Validate(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	transaction.UserID = authController.FetchLoggedInUserID(&c)
	err = transactionsServices.UpdateTransaction(&transaction, transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"transaction": transaction,
	})
}

func GetTransactions(c echo.Context) error {
	paginationData := pagination.ExtractPaginationData(&c)
	userID := authController.FetchLoggedInUserID(&c)
	strTime := c.QueryParam("month")
	month, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Invalid time format",
		})
	}
	transactions, numberOfRecords := transactionsServices.GetTransactionsByUser(paginationData, month, userID)
	return c.JSON(http.StatusOK, echo.Map{
		"transactions":    transactions,
		"numberOfRecords": numberOfRecords,
	})
}

func GetGroupedTransactions(c echo.Context) error {
	loggedInUserID := authController.FetchLoggedInUserID(&c)
	strTime := c.QueryParam("month")
	month, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Invalid time format",
		})
	}
	groupedTransactions := transactionsServices.GetGroupedTransactions(loggedInUserID, month)
	return c.JSON(http.StatusOK, echo.Map{
		"groupedTransactions": groupedTransactions,
	})
}

func DeleteTransaction(c echo.Context) error {
	transactionID, err := strconv.Atoi(c.Param("transactionID"))
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{})
	}
	loggedInUserID := authController.FetchLoggedInUserID(&c)
	err = transactionsServices.DeleteTransaction(uint(transactionID), loggedInUserID)
	if err != nil {
		if errors.Is(err, &customDatabaseErrors.RecordNotFoundError{}) {
			return c.JSON(http.StatusNotFound, echo.Map{})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	return c.JSON(http.StatusNoContent, echo.Map{})
}
