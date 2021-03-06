package transactions

import (
	customDatabaseErrors "backend/database/errors"
	authController "backend/middlewares/auth"
	categoriesModel "backend/models/transactions"
	categoriesServices "backend/services/transactions"
	utilsErrors "backend/utils/errors"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateCategory(c echo.Context) error {
	var category categoriesModel.Category
	c.Bind(&category)
	err := c.Validate(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	category.UserID = authController.FetchLoggedInUserID(&c)
	err = categoriesServices.CreateCategory(&category)
	if err != nil {
		if errors.Is(err, &utilsErrors.MaximumCategoriesError{}) {
			return c.JSON(http.StatusNotAcceptable, echo.Map{
				"message": err.Error(),
			})
		}
		if errors.Is(err, &customDatabaseErrors.DuplicateCategoryError{}) {
			return c.JSON(http.StatusConflict, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{})
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
	return c.NoContent(http.StatusNoContent)
}

func GetCategories(c echo.Context) error {
	loggedInUserID := authController.FetchLoggedInUserID(&c)
	categories := categoriesServices.GetCategories(loggedInUserID)
	return c.JSON(http.StatusOK, echo.Map{
		"categories": categories,
	})
}

func UpdateCategory(c echo.Context) error {
	var category categoriesModel.Category
	c.Bind(&category)
	categoryID := c.Param("categoryID")
	err := c.Validate(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	category.UserID = authController.FetchLoggedInUserID(&c)
	err = categoriesServices.UpdateCategory(&category, categoryID)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"category": category,
	})
}