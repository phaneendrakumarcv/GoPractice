package main

import (
	"fmt"
)

var accountBalance float64 = 1000.0

func main() {
	choice := 0
	fmt.Println("Welcome to Phantom Bank")
	fmt.Println("We offer following Services")
	fmt.Println(`1. View Balance
2. Debit Amount
3. Credit Amount
4. Good Bye
Select your choice: `)
	fmt.Scanf("%d", &choice)
	if choice == 1 {
		balance := viewBalance()
		fmt.Println("Your account balance is:", balance)
	} else if choice == 2 {
		var debitAmount = 0.0
		fmt.Println("Enter your debit amount: ")
		fmt.Scanf("%f", &debitAmount)
		accountBalance = debitBalance(debitAmount)
		fmt.Println("Your account balance is:", accountBalance)
	} else if choice == 3 {
		var creditAmount = 0.0
		fmt.Println("Enter your debit amount: ")
		fmt.Scanf("%f", &creditAmount)
		accountBalance = creditBalance(creditAmount)
		fmt.Println("Your account balance is:", accountBalance)
	} else {
		fmt.Println("Good bye")
	}
}

func debitBalance(amount float64) float64 {
	if amount <= accountBalance {
		accountBalance = accountBalance - amount
	} else {
		fmt.Println("Balance not enough")
	}
	return accountBalance
}

func creditBalance(amount float64) float64 {
	if amount > 0 {
		accountBalance = accountBalance + amount
	} else {
		fmt.Println("Please enter amount greater than 0")
	}
	return accountBalance
}

func viewBalance() float64 {
	return accountBalance
}
