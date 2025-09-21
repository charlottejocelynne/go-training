package main

import (
	"fmt"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func square(side int) {
	for i := 0; i < side; i++ {
		// tampilkan mendatar ke kanan......
		for j := 0; j < side; j++ {
			fmt.Print("* ")
		}
		// pindah baris
		fmt.Println()

	}
}

func rightTriangle(height int) {
	for i := 0; i <= height; i++ {
		// tampilkan mendatar ke kanan......
		for j := 0; j <= i-1; j++ {
			fmt.Print("* ")
		}
		// pindah baris
		fmt.Println()
	}

}

func upsideDownT(height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < height-i; j++ {
			fmt.Print("* ")
		}
		println()
	}
}

func equilateral(height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < height-i-1; j++ {
			fmt.Print("  ")
		}
		for j := 0; j < 1+2*i; j++ {
			fmt.Print("* ")
		}
		println()
	}
}

func upsideDownE(height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("  ")
		}
		for j := 0; j < 2*(height-i)-1; j++ {
			fmt.Print("* ")
		}
		println()
	}
}

func main() {
	fmt.Println("Soal 1")
	square(10)
	println()
	fmt.Print("Soal 2")
	rightTriangle(5)
	fmt.Println("Soal 3")
	upsideDownT(9)
	fmt.Println("Soal 4")
	equilateral(5)
	fmt.Println("Soal 5")
	upsideDownE(5)

}
