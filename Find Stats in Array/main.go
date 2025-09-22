package main

import "fmt"

func smallestNumber(arr []int) int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < temp {
			temp = arr[i]
		}
	}
	return temp
}

func biggestNumber(arr []int) int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] >= temp {
			temp = arr[i]
		}
	}
	return temp
}

func sum(arr []int) int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		x := temp + arr[i]
		temp = x
	}
	return temp
}

func avg(arr []int) float64 {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		x := temp + arr[i]
		temp = x
	}
	avg := float64(temp) / float64(len(arr))
	return avg
}

func main() {
	array := []int{3, 5, 8, 6, 2}
	smalest := smallestNumber(array)
	fmt.Println("The smallest number is ", smalest)
	biggest := biggestNumber(array)
	fmt.Println("The biggest number is ", biggest)
	sum := sum(array)
	fmt.Println("The sum is ", sum)
	avg := avg(array)
	fmt.Println("The average is ", avg)
}
