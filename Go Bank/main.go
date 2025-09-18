package main

import (
	"fmt"
	"strings"
	"time"
)

// Transaction declares a custom data structure for storing transaction history
type Transaction struct {
	Type    string  // "Deposit" atau "Withdraw"
	Amount  float64 // jumlah transaksi
	Balance float64 // saldo setelah transaksi
	Now     string
}

var balance float64

// history is an array of Transaction
var history []Transaction

func addTransaction(tType string, amount float64) {
	// make sure that amount is bigger than 0
	if amount < 0 {
		fmt.Println("‼️ Invalid ammount. Must be greater than 0 ‼️")
		return
	}

	// check if type if "D" (deposit), then add amount to balance
	if tType == "D" {
		balance += amount
		fmt.Println("Balance Updated! \nNew Amount: Rp", formatCurrency(balance))
	} else if tType == "W" {
		if amount <= balance {
			balance -= amount
			fmt.Println("Balance Updated! \nNew Amount:Rp", formatCurrency(balance))
		} else {
			fmt.Printf(" ‼️ ⚠️ Your balance is insufficient ⚠️ ‼️\nBalance: Rp%v\n", balance)
			return
		}
	}

	history = append(history, Transaction{
		Type:    tType,
		Amount:  amount,
		Balance: balance,
		Now:     time.Now().Format("2006-01-02 15:04:05"),
	})
}

// showHistory to show transaction history
func showHistory() {
	if len(history) != 0 {
		fmt.Println("=== Mutation History ===")
		for _, h := range history {
			fmt.Printf("%s | %s: Rp%s | Balance: Rp%s \n", h.Now, h.Type, formatCurrency(h.Amount), formatCurrency(h.Balance))
		}
	} else {
		fmt.Println("No Transaction History")
	}
}

// formatCurrency returns the given number, formated as currency
func formatCurrency(n float64) string {
	s := fmt.Sprintf("%.2f", n)

	// pisahkan integer dan desimal
	parts := strings.Split(s, ".")
	intPart := parts[0]
	decPart := parts[1]

	// kasih koma ribuan di intPart
	var result []string
	for i, c := range intPart {
		if (len(intPart)-i)%3 == 0 && i != 0 {
			result = append(result, ",")
		}
		result = append(result, string(c))
	}

	return strings.Join(result, "") + "." + decPart
}

func main() {
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
			fmt.Printf("Your balance is: Rp%v \n", formatCurrency(balance))
			showHistory()

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
			addTransaction("D", deposit)

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
			addTransaction("W", withdraw)

			// head back to the main menu
			continue
		default:
			fmt.Println("Goodbye!👋🏻")
		}
		break
	}
}
