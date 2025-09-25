package main

import (
	"fmt"
)

func increase(number *int, amount int) {
	*number += amount
}

func decrease(number *int, amount int) {
	*number -= amount
}

func main() {
	x := 10
	amount := 3
	number := &x
	fmt.Println("Before call: X = ", *number)
	increase(number, amount)
	fmt.Println("Increase X by 3: X = ", *number)
	decrease(number, amount)
	fmt.Println("Decrease X by 3: X = ", *number)
}
