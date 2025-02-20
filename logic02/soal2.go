package logic02

import "fmt"

func Soal2() {
	var matrix [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			matrix[i][j] = 2*j + 2
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
