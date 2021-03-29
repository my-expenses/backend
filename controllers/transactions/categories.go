package transactions

import (
	customDatabaseErrors "backend/database/errors"
	authController "backend/middlewares/auth"
	categoriesModel "backend/models/transactions"
	categoriesServices "backend/services/transactions"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateCategory(c echo.Context) error {
	var category categoriesModel.Category
	c.Bind(&category)
	category.UserID = authController.FetchLoggedInUserID(&c)
	err := categoriesServices.CreateCategory(&category)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"category": category,
	})
}

func DeleteCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{})
	}
	loggedInUserID := authController.FetchLoggedInUserID(&c)
	err = categoriesServices.DeleteCategory(uint(categoryID), loggedInUserID)
	if err != nil {
		if errors.Is(err, &customDatabaseErrors.RecordNotFoundError{}) {
			return c.JSON(http.StatusNotFound, echo.Map{})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	return c.JSON(http.StatusNoContent, echo.Map{})
}