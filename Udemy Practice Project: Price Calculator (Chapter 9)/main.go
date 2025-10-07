package main

import (
	"fmt"
	"main/filemanager"
	"main/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}

//	func getInputData() (prices []float64) {
//		fmt.Println("Enter 3 price")
//		var price1 float64
//		fmt.Scanln(&price1)
//		var price2 float64
//		fmt.Scanln(&price2)
//		var price3 float64
//		fmt.Scanln(&price3)
//		return []float64{price1, price2, price3}
//	}
//func priceCalculator(prices []float64, taxRate []float64) {
//	fmt.Printf("Tax Rate: ")
//	for _, rate := range taxRate {
//		fmt.Printf("%6.0f%%", rate)
//	}
//	fmt.Println()
//
//	for _, price := range prices {
//		fmt.Printf("Price %.0f:", price)
//		for _, rate := range taxRate {
//			final := price + (price * rate / 100)
//			fmt.Printf("%8.2f", final)
//		}
//		fmt.Println()
//	}
//}
