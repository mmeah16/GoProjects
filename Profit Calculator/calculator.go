package main

import (
	"fmt"
	"os"
)

// Goals
// 1. Validate user input
//    - show error message and exit if invalid input is provided
//	  - no negative numbers
//    - not 0
// 2. Store calculated results into file

const profitFile = "profile.txt"

func writeProfitToFile(ebt, profit string) {
	results := fmt.Sprintf("EBT: %s\nProfit: %s", ebt, profit)
	os.WriteFile(profitFile, []byte(results), 0644)

}

func main() {

	revenue, expenses, taxRate := getData()

	ebt, profit := calculateProfits(revenue, expenses, taxRate)
	writeProfitToFile(ebt, profit)
	fmt.Println(ebt)
	fmt.Println(profit)
}

func getData() (revenue, expenses, taxRate float64) {

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	if revenue <= 0 {
		panic("Revenue cannot be less than or equal to 0")
	} else if expenses < 0 {
		panic("Expenses cannot be negative!")
	} else if taxRate <= 0 {
		panic("Tax rate cannot be less than or equal to 0%")
	} else {
		return revenue, expenses, taxRate
	}

}

func calculateProfits(revenue, expenses, taxRate float64) (ebtString, profitString string) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	ebtString = fmt.Sprintf("Earnings before taxes: %.2f", ebt)
	profitString = fmt.Sprintf("Profits after taxes: %.2f", profit)

	return ebtString, profitString
}
