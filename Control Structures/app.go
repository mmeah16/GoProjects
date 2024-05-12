package main

import "fmt"

func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("How can we help?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1 , stores conditional in wantsCheckBalance as boolean

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Desposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	}
}
