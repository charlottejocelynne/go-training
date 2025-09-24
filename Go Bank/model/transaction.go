package model

// Transaction declares a custom data structure for storing transaction history
type Transaction struct {
	Type    string  // "Deposit" atau "Withdraw"
	Amount  float64 // jumlah transaksi
	Balance float64 // saldo setelah transaksi
	Now     string
}
