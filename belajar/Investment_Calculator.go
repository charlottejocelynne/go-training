package main

import "math"
import "fmt"

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64

	fmt.Print("Investment Amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years:")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = investmentAmount * math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)
}
