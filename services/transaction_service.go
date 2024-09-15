package services

import (
	"fmt"
	"splitwise/constants"
	"splitwise/models"
	"sync"
)

type TransactionService interface {
	CreateExpense(expenseType constants.ExpenseType, expense *models.Expense)
	GetBalance(userId string)
}

var transactionServiceInstance TransactionService
var transactionServiceSingleton sync.Once

type TransactionServiceImpl struct {
	balanceSheet   map[string]map[string]float64
	expenseService *ExpenseService
}

func GetTransactionService() TransactionService {
	transactionServiceSingleton.Do(func() {
		transactionServiceInstance = &TransactionServiceImpl{
			balanceSheet:   map[string]map[string]float64{},
			expenseService: NewExpenseService(),
		}
	})

	return transactionServiceInstance
}

func (s *TransactionServiceImpl) GetBalance(userId string) {
	for balanceUser, balanceAmount := range s.balanceSheet[userId] {
		if balanceAmount < 0 {
			fmt.Println("%s owed by %s: %0.2f", balanceUser, userId, balanceAmount)
		} else {
			fmt.Println("%s owes %s: %0.2f", balanceUser, userId, balanceAmount)
		}
	}
}

func (s *TransactionServiceImpl) CreateExpense(expenseType constants.ExpenseType, expense *models.Expense) {
	splits := expense.Splits
	expense, _ = s.expenseService.ComputeExpenseSplit(expenseType, expense)
	fmt.Println(expense)

	for _, split := range splits {
		if paidByUserBalance, ok := s.balanceSheet[expense.PaidBy.Id]; ok {
			if _, ok := paidByUserBalance[split.UserId]; ok {
				s.balanceSheet[expense.PaidBy.Id][split.UserId] += split.Amount
			} else {
				s.balanceSheet[expense.PaidBy.Id][split.UserId] = split.Amount
			}
		} else {
			s.balanceSheet[expense.PaidBy.Id] = map[string]float64{split.UserId: split.Amount}
		}

		if paidByUserBalance, ok := s.balanceSheet[split.UserId]; ok {
			if _, ok := paidByUserBalance[expense.PaidBy.Id]; ok {
				s.balanceSheet[split.UserId][expense.PaidBy.Id] -= split.Amount
			} else {
				s.balanceSheet[split.UserId][expense.PaidBy.Id] = -split.Amount
			}
		} else {
			s.balanceSheet[split.UserId] = map[string]float64{expense.PaidBy.Id: -split.Amount}
		}
	}
}
