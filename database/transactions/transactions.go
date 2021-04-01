package transactions

import (
	dbInstance "backend/database"
	customDatabaseErrors "backend/database/errors"
	transactionsModel "backend/models/transactions"
)

func CreateTransaction(transaction *transactionsModel.Transaction) error {
	return dbInstance.GetDBConnection().Create(transaction).Error
}

func DeleteTransaction(transactionID, loggedInUserID uint) error {
	db := dbInstance.GetDBConnection().Unscoped().Where("id = ? AND user_id = ?",
		transactionID, loggedInUserID).Delete(&transactionsModel.Transaction{})
	if db.RowsAffected == 0 {
		return &customDatabaseErrors.RecordNotFoundError{}
	}
	return db.Error
}