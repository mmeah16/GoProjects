package main

import "fmt"

func main() {

	var revenue = 0.0
	var expenses = 0.0
	var taxRate = 0.0

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := (revenue - expenses)
	profit := ebt * (1 - taxRate/100)
	fmt.Println(ebt)
	fmt.Println(profit)
}
