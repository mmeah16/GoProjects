package main

import (
	"fmt"
	"math"
)

// scoped to all functions
// const cannot be changed or overwritten, yet var can
const inflationRate = 2.5

func main() {
	// overriding inferred value
	var investmentAmount float64
	var years float64
	// declare and assign variable where type can be inferred by Go
	expectedReturnRate := 0.0

	// & - pointer to a variable, allows scan to populate variable with a value
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Investment Length in Years: ")
	fmt.Scan(&years)
	fmt.Print("Investment Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+(expectedReturnRate)/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Sprintf saves formatted string in a variable
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Println("Future Value:",futureValue)
	// fmt.Printf("Future Value: %.1f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+(expectedReturnRate)/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	// explicitly state what you want to return
	return fv, rfv

	// automatically returns the values the function generates
	// return
}
