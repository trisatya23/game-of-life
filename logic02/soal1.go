package logic02

import "fmt"

func Soal1() {
	var matrix [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			matrix[i][j] = 2*j + 1
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
