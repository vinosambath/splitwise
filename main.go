package main

import (
	"fmt"
	"splitwise/constants"
	"splitwise/handlers"
	"splitwise/models"
	"splitwise/models/split"
)

/*

---------
Models

Users -> id, name,
Expense -> userId, paid by, splits
***** Future Scoping: TransactionHistory -> id, payee, payer, amount

---------

Split types

Splits -> payee, amount

---------

Services

UserService -> getUser, createUser,
Expense ->
   map[payee][payer]float64
   ***** Future Scoping:
        addTransaction -> update balance of each user to user pair, create transaction history for listing transactions of one particular user
***** Future Scoping: TransactionHistory - map[userid] = list[transactions]

---------

Handler

ShowBalance
CreateExpense

---------

*/

func main() {
	fmt.Println("Hello")

	user1 := models.User{
		Id:   "1",
		Name: "name1",
	}
	user2 := models.User{
		Id:   "2",
		Name: "name2",
	}
	user3 := models.User{
		Id:   "3",
		Name: "name3",
	}

	handler := handlers.NewExpenseHandler()
	handler.CreateExpense(constants.ExpenseTypePercent, &models.Expense{
		PaidBy: user1,
		Amount: 100,
		Splits: []*split.Split{
			{
				UserId:  user2.Id,
				Percent: 50,
			},
			{
				UserId:  user3.Id,
				Percent: 50,
			},
		},
	})

	handler.CreateExpense(constants.ExpenseTypeExact, &models.Expense{
		PaidBy: user2,
		Amount: 50,
		Splits: []*split.Split{
			{
				UserId:  user1.Id,
				Percent: 10,
			},
			{
				UserId:  user3.Id,
				Percent: 40,
			},
		},
	})

	handler.GetBalance(user1.Id)
	fmt.Println("second user balance ")
	handler.GetBalance(user2.Id)
	fmt.Println("third user balance ")
	handler.GetBalance(user3.Id)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
