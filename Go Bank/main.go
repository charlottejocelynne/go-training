package main

import (
	"fmt"
	"gobank/model"
	"gobank/service"
	"gobank/utils"
	"time"
)

func main() {
	var balance float64
	var history []model.Transaction
	// perform infinite loop, until break by the exit menu
	for {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04 WITA"))
		fmt.Println("˚.🎀༘⋆Welcome to the Go Bank!˚.🎀༘⋆")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Balance and Mutation")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Your Choice:")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// show the instructions
			fmt.Println("Your choice is 1. Balance and Mutation")
			fmt.Printf("Your balance is: Rp%v \n", utils.FormatCurrency(balance))
			service.ShowHistory(history)

			// head back to the main menu
			fmt.Println("Press enter to continue")
			fmt.Scanln()
			continue
		case 2:
			// show the instructions
			fmt.Println("Your choice is 2. Deposit")
			fmt.Printf("Deposit:")

			// request for withdrawal amount
			var deposit float64
			fmt.Scan(&deposit)

			// perform the transaction
			balance, history = service.AddTransaction(balance, "D", deposit, history)

			// head back to the main menu
			fmt.Println("Press enter to continue")
			fmt.Scanln()
			continue
		case 3:
			// show the instructions
			fmt.Println("Your choice is 3")
			fmt.Printf("Withdraw :")

			// request for withdrawal amount
			var withdraw float64
			fmt.Scan(&withdraw)

			// perform the transaction
			service.AddTransaction(balance, "W", withdraw, history)

			// head back to the main menu
			continue
		default:
			fmt.Println("Goodbye!👋🏻")
		}
		break
	}
}
