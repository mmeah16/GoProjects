package main

import (
	"fmt"
	"os"
	"strconv"
)

// write data to a file

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

	for i := 0; i < 200; i++ {

		fmt.Println("Welcome to Go Bank!")
		fmt.Println("How can we help?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
			fmt.Println()
		case 2:
			fmt.Print("Desposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than $0.00")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
			fmt.Println()
			fmt.Println()
		case 3:
			fmt.Print("Withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Withdrawl amount must be greater than $0.00")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Withdrawl amount is greater than account balance!")
				continue
			}
			accountBalance -= withdrawlAmount
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for visiting Go Bank!")
			fmt.Println()

			// return must be used to end the entire application
			// switch should be used when application does not need to be broken out of
			return
		}

		// wantsCheckBalance := choice == 1 , stores conditional in wantsCheckBalance as boolean

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Desposit amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Deposit amount must be greater than $0.00")
		// 		return
		// 	}
		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdrawl amount: ")
		// 	var withdrawlAmount float64
		// 	fmt.Scan(&withdrawlAmount)

		// 	if withdrawlAmount <= 0 {
		// 		fmt.Println("Withdrawl amount must be greater than $0.00")
		// 		return
		// 	}

		// 	if withdrawlAmount > accountBalance {
		// 		fmt.Println("Withdrawl amount is greater than account balance!")
		// 		return
		// 	}
		// 	accountBalance -= withdrawlAmount
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}

}
