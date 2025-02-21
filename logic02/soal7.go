package logic02

import "fmt"

func Soal7(n int) {
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("    ")
		}
		fmt.Println(start)
		start += 2
	}
}
