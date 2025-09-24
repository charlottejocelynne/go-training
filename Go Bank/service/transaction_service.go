package service

import (
	"fmt"
	"gobank/model"
	"gobank/utils"
	"time"
)

//var Balance float64
//var history []model.Transaction

func AddTransaction(
	balance float64,
	tType string,
	amount float64,
	history []model.Transaction,
) (float64, []model.Transaction) {
	// make sure that amount is bigger than 0
	if amount < 0 {
		fmt.Println("‼️ Invalid ammount. Must be greater than 0 ‼️")
		return balance, history
	}

	// check if type if "D" (deposit), then add amount to Balance
	if tType == "D" {
		balance += amount
		fmt.Println("Balance Updated! \nNew Amount: Rp", utils.FormatCurrency(balance))
	} else if tType == "W" {
		if amount <= balance {
			balance -= amount
			fmt.Println("Balance Updated! \nNew Amount:Rp", utils.FormatCurrency(balance))
		} else {
			fmt.Printf(" ‼️ ⚠️ Your Balance is insufficient ⚠️ ‼️\nBalance: Rp%v\n", balance)
			return balance, history
		}
	}
	history = append(history, model.Transaction{
		Type:    tType,
		Amount:  amount,
		Balance: balance,
		Now:     time.Now().Format("2006-01-02 15:04:05"),
	})
	return balance, history
}

func ShowHistory(history []model.Transaction) []model.Transaction {
	if len(history) != 0 {
		fmt.Println("=== Mutation History ===")
		for _, h := range history {
			fmt.Printf("%s | %s: Rp%s | Balance: Rp%s \n", h.Now, h.Type, utils.FormatCurrency(h.Amount), utils.FormatCurrency(h.Balance))
		}
	} else {
		fmt.Println("No Transaction History")
	}
	return history
}
