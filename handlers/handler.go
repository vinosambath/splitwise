package handlers

import (
	"splitwise/constants"
	"splitwise/models"
	"splitwise/services"
)

type ExpenseHandler struct {
	transactionService services.TransactionService
}

func NewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{
		services.GetTransactionService(),
	}
}

func (h *ExpenseHandler) CreateExpense(expenseType constants.ExpenseType, expense *models.Expense) {
	h.transactionService.CreateExpense(expenseType, expense)
}

func (h *ExpenseHandler) GetBalance(userId string) {
	h.transactionService.GetBalance(userId)
}
