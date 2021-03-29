package transactions

import (
	dbInstance "backend/database"
	customDatabaseErrors "backend/database/errors"
	categoriesModel "backend/models/transactions"
)

func CreateCategory(category *categoriesModel.Category) error {
	return dbInstance.GetDBConnection().Create(category).Error
}

func DeleteCategory(categoryID, loggedInUserID uint) error {
	db := dbInstance.GetDBConnection().Unscoped().Where("id = ? AND user_id = ?",
		categoryID, loggedInUserID).Delete(&categoriesModel.Category{})
	if db.RowsAffected == 0 {
		return &customDatabaseErrors.RecordNotFoundError{}
	}
	return db.Error
}