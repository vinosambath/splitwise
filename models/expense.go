package models

import "splitwise/models/split"

type Expense struct {
	PaidBy User
	Amount float64
	Splits []*split.Split
}
