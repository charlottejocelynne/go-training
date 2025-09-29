package model

import (
	"time"
)

type Book struct {
	Date            time.Time
	TransactionType string
	Amount          float64
	Balance         float64
}

type Books []Book
