package transactions

import (
	transactionsDBInteractions "backend/database/transactions"
	transactionsModel "backend/models/transactions"
)

func CreateTransaction(transaction *transactionsModel.Transaction) error {
	return transactionsDBInteractions.CreateTransaction(transaction)
}

func DeleteTransaction(transactionID, userID uint) error {
	return transactionsDBInteractions.DeleteTransaction(transactionID, userID)
}