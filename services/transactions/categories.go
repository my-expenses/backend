package transactions

import (
	categoriesDBInteractions "backend/database/transactions"
	categoriesModel "backend/models/transactions"
	utilsErrors "backend/utils/errors"
	"strconv"
)

func CreateCategory(category *categoriesModel.Category) error {
	categoriesCount := categoriesDBInteractions.GetCategoriesCount(category.UserID)
	if categoriesCount >= 10 {
		return &utilsErrors.MaximumCategoriesError{}
	}
	return categoriesDBInteractions.CreateCategory(category)
}

func DeleteCategory(categoryID, loggedInUserID uint) error {
	return categoriesDBInteractions.DeleteCategory(categoryID, loggedInUserID)
}

func GetCategories(userID uint) []categoriesModel.Category {
	return categoriesDBInteractions.GetCategories(userID)
}

func UpdateCategory(category *categoriesModel.Category, categoryID string) error {
	uintID, _ := strconv.ParseUint(categoryID, 10, 64)
	category.ID = uint(uintID)
	return categoriesDBInteractions.UpdateCategory(category)
}