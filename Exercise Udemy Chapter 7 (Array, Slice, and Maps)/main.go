package main

import "fmt"

type productList struct {
	title string
	id    string
	price float64
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobby := [3]string{"swim", "music", "movie"}
	fmt.Println(hobby)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobby[0])
	fmt.Println(hobby[1:])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	Hobby2 := hobby[:2]
	fmt.Println(Hobby2)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	Hobby2 = Hobby2[1:3]
	fmt.Println(Hobby2)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Build portfolio", "Gain experience"}
	fmt.Println(goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Learn Programming"
	fmt.Println(goals)
	goals = append(goals, "Gain experience")
	fmt.Println(goals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []productList{
		{
			"first product",
			"a first product",
			12.99,
		},
		{
			"second product",
			"a second product",
			129.99,
		},
	}
	fmt.Println(products)
	newProducts := productList{
		"third product",
		"a third product",
		15.99,
	}
	products = append(products, newProducts)
}
