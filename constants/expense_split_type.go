package constants

type ExpenseType int

const (
	ExpenseTypeEqual   ExpenseType = iota
	ExpenseTypeExact   ExpenseType = iota
	ExpenseTypePercent ExpenseType = iota
)
