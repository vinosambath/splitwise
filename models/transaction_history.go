package models

type TransactionHistory struct {
	UserId    string
	Amount    float64
	OtherUser User
}
