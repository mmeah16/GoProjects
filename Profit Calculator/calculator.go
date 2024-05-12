package main

import "fmt"

func main() {

	revenue, expenses, taxRate := getData()

	ebt, profit := calculateProfits(revenue, expenses, taxRate)

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

	return revenue, expenses, taxRate
}

func calculateProfits(revenue, expenses, taxRate float64) (ebtString, profitString string) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	ebtString = fmt.Sprintf("Earnings before taxes: %.2f", ebt)
	profitString = fmt.Sprintf("Profits after taxes: %.2f", profit)

	return ebtString, profitString
}
