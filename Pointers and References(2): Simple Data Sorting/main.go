package main

import (
	"fmt"
)

func sortAsc(arr *[]int) {
	if len(*arr) == 0 {
	}
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			// save the i-th element into a variable named temp
			temp := (*arr)[i]
			if temp > (*arr)[j] {
				// return the i-th element into temp variable
				temp = (*arr)[i]
				// return i-th element to j-th element
				(*arr)[i] = (*arr)[j]
				// assign the previous value back to j-th element
				(*arr)[j] = temp
			}

		}
	}
}

func main() {
	x := []int{5, 3, 8, 6, 2, -1, 10, 11}
	array := &x
	fmt.Println("Original Array:", *array)
	sortAsc(array)
	fmt.Println("Sorted Ascending", *array)
}
