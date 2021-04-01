package transactions

import (
	dbInstance "backend/database"
	customDatabaseErrors "backend/database/errors"
	categoriesModel "backend/models/transactions"
)

func GetCategoriesCount(userID uint) int64 {
	categoriesCount := int64(0)
	dbInstance.GetDBConnection().Model(&categoriesModel.Category{}).Where("user_id = ?", userID).Count(&categoriesCount)
	return categoriesCount
}

func CreateCategory(category *categoriesModel.Category) error {
	db := dbInstance.GetDBConnection().Create(category)
	if db.RowsAffected == 0 {
		return &customDatabaseErrors.DuplicateCategoryError{}
	}
	return db.Error
}

func DeleteCategory(categoryID, loggedInUserID uint) error {
	db := dbInstance.GetDBConnection().Unscoped().Where("id = ? AND user_id = ?",
		categoryID, loggedInUserID).Delete(&categoriesModel.Category{})
	if db.RowsAffected == 0 {
		return &customDatabaseErrors.RecordNotFoundError{}
	}
	return db.Error
}

func GetCategories(userID uint) []categoriesModel.Category {
	categories := make([]categoriesModel.Category, 0)
	dbInstance.GetDBConnection().Where("user_id = ?", userID).Order("created_at").Find(&categories)
	return categories
}

func UpdateCategory(category *categoriesModel.Category) error {
	db := dbInstance.GetDBConnection().Model(category).Update("title", category.Title)
	if db.RowsAffected == 0 {
		return &customDatabaseErrors.DuplicateCategoryError{}
	}
	return db.Error
}