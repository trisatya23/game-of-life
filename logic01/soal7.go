package logic01

import "fmt"

func Soal7() {
	angka := 1
	n := 10
	median := 0.5 * float64(n) //3
	for i := 1; i <= n; i++ {  //5
		fmt.Println(angka, " ")  //3
		if float64(i) < median { //5<3 false
			angka = angka + 2 //3+2 = 5
		} else if float64(i) > median { //5>3 true
			angka = angka - 2 //3-2 = 1
		}
	}
}
