package transactions

import (
	customDatabaseErrors "backend/database/errors"
	authController "backend/middlewares/auth"
	transactionsModel "backend/models/transactions"
	transactionsServices "backend/services/transactions"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
