package services

import (
	"splitwise/models"
	"sync"
)

type TransactionHistoryService interface {
	AddTransactionHistory(user models.User, expense *models.Expense)
	GetTransactionHistory(userId string) ([]*models.TransactionHistory, error)
}

type TransactionHistoryImpl struct {
	inMemTransactionHistory map[string][]*models.TransactionHistory
}

var TransactionHistoryInstance TransactionHistoryService
var TransactionHistoryInstanceOnce sync.Once

func NewTransactionHistoryService() TransactionHistoryService {
	TransactionHistoryInstanceOnce.Do(func() {
		TransactionHistoryInstance = &TransactionHistoryImpl{}
	})
	return TransactionHistoryInstance
}

func (t *TransactionHistoryImpl) AddTransactionHistory(user models.User, expense *models.Expense) {
	transactionHistory := &models.TransactionHistory{
		UserId: user.Id,
		Amount: expense.Amount,
	}

	t.inMemTransactionHistory[user.Id] = append(t.inMemTransactionHistory[user.Id], transactionHistory)

	for _, split := range expense.Splits {
		transactionHistory := &models.TransactionHistory{
			UserId: split.UserId,
			Amount: split.Amount,
		}
		t.inMemTransactionHistory[split.UserId] = append(t.inMemTransactionHistory[split.UserId], transactionHistory)
	}
}

func (t *TransactionHistoryImpl) GetTransactionHistory(userId string) (transactionHistory []*models.TransactionHistory, err error) {
	return t.inMemTransactionHistory[userId], nil
}
