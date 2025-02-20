package logic02

import "fmt"

func Soal05() {
	n := 9
	matrix := make([][]int, n)
	start := 1

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)

		if i%2 == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = start
				start += 3
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				matrix[i][j] = start
				start += 3
			}
		}
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
}
