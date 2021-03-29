package transactions

import (
	dbInstance "backend/database"
	categoriesModel "backend/models/transactions"
)

func CreateCategory(category *categoriesModel.Category) error {
	return dbInstance.GetDBConnection().Create(category).Error
}