package main

import (
	"fmt"
	"strings"
	"time"
)

func accountNumber() int {
	accountNumber := 101010102003
	println("Account Number:", accountNumber)
	return accountNumber
}
func formating(n float64) string {
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

func saveBalance(balance float64) float64 {
	return balance
}

var history []Transaction

type Transaction struct {
	Type    string  // "Deposit" atau "Withdraw"
	Amount  float64 // jumlah transaksi
	Balance float64 // saldo setelah transaksi
	Now     string
}

func addTransaction(tType string, amount float64, balance float64) {
	history = append(history, Transaction{
		Type:    tType,
		Amount:  amount,
		Balance: balance,
		Now:     time.Now().Format("2006-01-02 15:04:05"),
	})
}

func showHistory() {
	if len(history) != 0 {
		fmt.Println("=== Mutation History ===")
		for _, h := range history {
			fmt.Printf("%s | %s: Rp%s | Balance: Rp%s \n", h.Now, h.Type, formating(h.Amount), formating(h.Balance))
		}
		return
	} else if len(history) == 0 {
		fmt.Println("No Transaction History")
	}
}

func main() {
	var balance float64
	for i := 0; i < 50; i++ {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04 WITA"))
		fmt.Println("˚.🎀༘⋆Welcome to the Go Bank!˚.🎀༘⋆")
		accountNumber()
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
			saveBalance(balance)
			fmt.Println("Your choice is 1. Balance and Mutation")
			fmt.Printf("Your balance is: Rp%v \n", formating(balance))
			showHistory()
			fmt.Println("Press enter to continue")
			fmt.Scanln()
			continue
		case 2:
			fmt.Println("Your choice is 2. Deposit")
			var deposit float64
			fmt.Printf("Deposit:")
			fmt.Scan(&deposit)
			if deposit > 0 {
				balance += deposit
				fmt.Println("Balance Updated! \nNew Amount: Rp", formating(balance))
				addTransaction("Deposit", deposit, balance)
			} else if deposit < 0 {
				fmt.Println("‼️ Invalid ammount. Must be greater than 0 ‼️")
			}
			continue
		case 3:
			fmt.Println("Your choice is 3")
			var withdraw float64
			fmt.Printf("Withdraw :")
			fmt.Scan(&withdraw)
			if withdraw < balance {
				balance -= withdraw
				fmt.Println("Balance Updated! \nNew Amount:Rp", formating(balance))
				addTransaction("Withdraw", withdraw, balance)
			} else if withdraw > balance {
				fmt.Printf(" ‼️ ⚠️ Your balance is insufficient ⚠️ ‼️\nBalance: Rp%v\n", balance)
			}
			continue
		default:
			fmt.Println("Goodbye!👋🏻")
		}
		break
	}
}
