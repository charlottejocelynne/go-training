package main

import (
	"fmt"
	"main/service"
	"main/utils"
	"time"
)

func main() {
	var balance float64
	history := service.Books{}

	for {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04 WITA"))
		fmt.Println("˚.🎀༘⋆Welcome to the Go Bank!˚.🎀༘⋆")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Balance and Mutation")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Your Choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your choice is 1. Balance and Mutation")
			fmt.Printf("Your balance is: Rp%v \n", utils.FormatCurrency(balance))
			history.ShowHistory()

		case 2:
			fmt.Println("Your choice is 2. Deposit")
			fmt.Print("Deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			t := service.Book{
				Date:            time.Now(),
				TransactionType: "D",
				Amount:          deposit,
			}
			t.AddTransaction(&balance, &history)

		case 3:
			fmt.Println("Your choice is 3. Withdraw")
			fmt.Print("Withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			t := service.Book{
				Date:            time.Now(),
				TransactionType: "W",
				Amount:          withdraw,
			}
			t.AddTransaction(&balance, &history)

		case 4:
			fmt.Println("Goodbye!👋🏻")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
