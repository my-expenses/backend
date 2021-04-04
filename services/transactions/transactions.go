package transactions

import (
	transactionsDBInteractions "backend/database/transactions"
	transactionsModel "backend/models/transactions"
)

func CreateTransaction(transaction *transactionsModel.Transaction) error {
	if *transaction.CategoryID == 0 {
		transaction.CategoryID = nil
	}
	return transactionsDBInteractions.CreateTransaction(transaction)
}

func DeleteTransaction(transactionID, userID uint) error {
	return transactionsDBInteractions.DeleteTransaction(transactionID, userID)
}