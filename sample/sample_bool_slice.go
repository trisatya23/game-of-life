package sample

import "fmt"

func SampleBoolSlice() {

	slice := [][]bool{{false, true, false}, {false, false, true},
		{false, true, false}, {false, true, true},
		{false, false, false}, {false, false, false}}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] {
				fmt.Print("0\t")
			} else {
				fmt.Print("-\t")
			}
		}
		fmt.Println()
	}
}
