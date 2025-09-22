package main

import "fmt"

func prime(x int) {
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

func oddEven(x int) {
	if x%2 == 0 {
		fmt.Print("Number is even ")
	} else {
		fmt.Print("Number is odd ")
	}
}

func main() {
	for {
		fmt.Print("Enter your number: ")
		var x int
		fmt.Scan(&x)
		oddEven(x)
		prime(x)
	}
}
