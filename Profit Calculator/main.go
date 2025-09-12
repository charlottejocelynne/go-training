package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	fmt.Print("How much your monthly revenue?")
	fmt.Scan(&revenue)
	fmt.Print("How much your monthly expenses?")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate:")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
