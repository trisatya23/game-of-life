package logic02

import "fmt"

func Soal03() {
	var matrix [9][9]int
	counter := 1

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			matrix[i][j] = counter
			counter += 2
		}
	}

	fmt.Println("Matriks:")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
