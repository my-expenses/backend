package transactions

import (
	authController "backend/middlewares/auth"
	categoriesModel "backend/models/transactions"
	categoriesServices "backend/services/transactions"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateCategory(c echo.Context) error {
	var category categoriesModel.Category
	c.Bind(&category)
	category.UserID = authController.FetchLoggedInUserID(c)
	err := categoriesServices.CreateCategory(&category)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"category": category,
	})
}