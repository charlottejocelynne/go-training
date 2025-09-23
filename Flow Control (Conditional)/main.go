package main

import "fmt"

func calcVolume(x, y, z int) int {
	return x * y * z
}

func cubeShape(x, y, z int) bool {
	return x == y && y == z && x == z
}

func main() {
	for {
		var x int
		fmt.Print("Input panjang: ")
		fmt.Scan(&x)
		var y int
		fmt.Print("Input lebar: ")
		fmt.Scan(&y)
		var z int
		fmt.Print("Input tinggi: ")
		fmt.Scan(&z)
		fmt.Println("=============")
		volume := calcVolume(x, y, z)
		if x > 0 && y > 0 && z > 0 {
			fmt.Println("Volume: ", volume)
			fmt.Print("Bangun ruang: ")
			if cubeShape(x, y, z) {
				fmt.Println("Kubus")
			} else {
				fmt.Println("Balok")
			}
		} else {
			fmt.Println("Bangun ruang: tidak diketahui!")
			fmt.Println("Volume: tidak dapat dihitung!")
		}
	}
}
