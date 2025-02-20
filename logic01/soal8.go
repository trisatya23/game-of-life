package logic01

import "fmt"

func Soal8() {
	angka := 2
	n := 10
	median := 0.5 * float64(n)

	for i := 1; i <= n; i++ {
		fmt.Println(angka, " ")
		if float64(i) < median {
			angka = angka + 2
		} else if float64(i) > median {
			angka = angka - 2
		}

	}
}
