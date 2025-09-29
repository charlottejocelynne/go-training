package service

import (
	"fmt"
	"main/utils"
	"time"
)

type Book struct {
	Date            time.Time
	TransactionType string
	Amount          float64
	Balance         float64
}

type Books []Book

func (t *Book) AddTransaction(balance *float64, history *Books) {
	if t.Amount < 0 {
		fmt.Println("‼️ Invalid amount. Must be greater than 0 ‼️")
		return
	}
	if t.TransactionType == "D" {
		*balance += t.Amount
		fmt.Println("Balance Updated! \nNew Amount: Rp", utils.FormatCurrency(*balance))
	} else if t.TransactionType == "W" {
		if t.Amount <= *balance {
			*balance -= t.Amount
			fmt.Println("Balance Updated! \nNew Amount: Rp", utils.FormatCurrency(*balance))
		} else {
			fmt.Printf("‼️ ⚠️ Your Balance is insufficient ⚠️ ‼️\nBalance: Rp%v\n", *balance)
			return
		}
	}

	t.Balance = *balance
	*history = append(*history, *t)
}

func (history *Books) ShowHistory() {
	if len(*history) != 0 {
		fmt.Println("=== Mutation History ===")
		for _, h := range *history {
			fmt.Printf("%s | %s: Rp%s | Balance: Rp%s \n",
				h.Date.Format("2006-01-02 15:04:05"),
				h.TransactionType,
				utils.FormatCurrency(h.Amount),
				utils.FormatCurrency(h.Balance))
		}
	} else {
		fmt.Println("No Transaction History")
	}
}
