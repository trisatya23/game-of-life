package logic01

import "fmt"

func Soal4() {
	n := 10
	angkapertama := 19

	for i := 0; i < n; i++ {
		fmt.Print(angkapertama-(2*i), " ")
	}
}
