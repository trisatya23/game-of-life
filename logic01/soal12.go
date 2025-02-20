package logic01

import "fmt"

func Soal12() {
	n := 12

	for i := 0; i < n/4; i++ {
		for j := 1; j <= 7; j += 2 {
			fmt.Println(j)
		}
	}
}
