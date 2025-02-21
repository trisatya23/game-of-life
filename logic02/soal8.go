package logic02

import "fmt"

func Soal8(n int) {
	start := 1

	number := make([]int, n)
	for i := 0; i < n; i++ {
		number[i] = start
		start += 2
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("    ")
		}
		fmt.Println(number[i])
	}
}
