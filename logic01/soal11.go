package logic01

import "fmt"

func Soal11() {
	num := 1
	n := 10

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(num)
			num *= 3
		}
	}

}
