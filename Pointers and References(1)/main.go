package main

import (
	"fmt"
)

func increase(numberPointers, amountPointers *int) {
	*numberPointers = *numberPointers + *amountPointers
}

func decrease(numberPointers, amountPointers *int) {
	*numberPointers = *numberPointers + *amountPointers
}

func main() {
	number := 10
	amount := 3
	numberPointers := &number
	amountPointers := &amount
	fmt.Println("Before call: X = ", *numberPointers)
	increase(numberPointers, amountPointers)
	fmt.Println("Increase X by 3: X = ", *numberPointers)
	decrease(numberPointers, amountPointers)
	fmt.Println("Decrease X by 3: X = ", *numberPointers)
}
