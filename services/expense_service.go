package services

import (
	"fmt"
	"splitwise/constants"
	"splitwise/models"
)

type ExpenseService struct{}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{}
}

func (exp *ExpenseService) ComputeExpenseSplit(expenseType constants.ExpenseType, expense *models.Expense) (*models.Expense, error) {

	if err := exp.validateExpenseSplit(expenseType, expense); err != nil {
		return nil, err
	}

	switch expenseType {
	case constants.ExpenseTypeEqual:
		equalExpense := expense.Amount / float64(len(expense.Splits))
		for _, split := range expense.Splits {
			split.Amount = equalExpense
		}
	case constants.ExpenseTypeExact:
		for _, split := range expense.Splits {
			split.Amount = expense.Amount
		}
	case constants.ExpenseTypePercent:
		for _, split := range expense.Splits {
			split.Amount = expense.Amount * float64(split.Percent) / 100
		}
	}
	return expense, nil
}

func (exp *ExpenseService) validateExpenseSplit(expenseType constants.ExpenseType, expense *models.Expense) error {
	if expenseType == constants.ExpenseTypeEqual || expenseType == constants.ExpenseTypeExact {
		return nil
	}
	// expense type is percent

	percentSum := (0)
	for _, split := range expense.Splits {
		percentSum += split.Percent
	}

	if percentSum != 100 {
		return fmt.Errorf("invalid input, expense split percent sum total is not 100")
	}

	return nil
}
