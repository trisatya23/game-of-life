package main

import (
	"fmt"
	"github.com/trisatya23/game-of-life/logic02"
)

func main() {
	n := 9

	result := logic02.Soal6c(n)
	fmt.Println("Matriks:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
