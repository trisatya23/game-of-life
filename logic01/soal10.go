package logic01

import "fmt"

func Soal10() {
	num := 2
	n := 10

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(num)
			num *= 2
		}
	}
}
