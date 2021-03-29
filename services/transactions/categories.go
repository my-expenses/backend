package transactions

import (
	categoriesDBInteractions "backend/database/transactions"
	categoriesModel "backend/models/transactions"
)

func CreateCategory(category *categoriesModel.Category) error {
	return categoriesDBInteractions.CreateCategory(category)
}

func DeleteCategory(categoryID, loggedInUserID uint) error {
	return categoriesDBInteractions.DeleteCategory(categoryID, loggedInUserID)
}