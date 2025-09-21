package main

import "fmt"

func main() {
	fmt.Print("Enter your number: ")
	var x int
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Print("Number is even ")
	} else {
		fmt.Print("Number is odd ")
	}
	if x < 2 {
		fmt.Println("and not a prime number")
		return
	}
	for i := x - 1; i >= 2; i-- {
		if x%i == 0 {
			fmt.Println("and not a prime number")
			return
		}
	}
	fmt.Println("and prime number")
}
