package logic02

import "fmt"

func Soal9(n int) {
	start := 1

	number := Create2DArray(n)

	for i := 0; i < n; i++ {
		number[i][i] = start
		number[n-i-1][i] = start
		start = start + 2
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(number[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}

}
