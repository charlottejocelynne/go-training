package main

import (
	"fmt"
)

func sortAsc(arr [5]int) []int {
	if len(arr) == 0 {
		return arr[0:5]
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			// save the i-th element into a variable named temp
			temp := arr[i]
			if temp > arr[j] {
				// return the i-th element into temp variable
				temp = arr[i]
				// return i-th element to j-th element
				arr[i] = arr[j]
				// return back the previous value back to j-th element
				arr[j] = temp
			}

		}
	}
	return arr[0:5]
}

func main() {
	x := [5]int{5, 3, 8, 6, 2}
	sorted := sortAsc(x)
	fmt.Println("Original Array:", x)
	fmt.Println("Sorted Ascending", sorted)

}
