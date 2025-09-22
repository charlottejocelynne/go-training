package main

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := x - 1; i >= 2; i-- {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isEven(x int) bool {
	//if x%2 == 0 {
	//	return true
	//}
	//return false
	return x%2 == 0
}

func main() {
	for {
		fmt.Print("Enter your number: ")
		var x int
		fmt.Scan(&x)
		//oddEven(x)
		//prime(x)
		if isEven(x) {
			fmt.Print("Number is even ")
		} else {
			fmt.Print("Number is odd ")
		}
		if isPrime(x) {
			fmt.Println("and prime number")
		} else {
			fmt.Println("and not a prime number")
		}
	}
}
