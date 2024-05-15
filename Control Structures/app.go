package main

import "fmt"

func main() {
	var accountBalance = 1000.00

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
				return
			}
			accountBalance -= withdrawlAmount
		default:
			fmt.Println("Goodbye!")
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
	fmt.Println("Thanks for visiting Go Bank!")
}
