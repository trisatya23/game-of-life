package sample

import "fmt"

func DinamicSlice(row, col int) {
	fmt.Println("Dinamic Slice")
	cell := make([][]int, row)
	for i := 0; i < len(cell); i++ {
		cell[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == j || i+j == row-1 {
				fmt.Print("#\t")
			} else {
				fmt.Print(".\t")
			}
		}
		fmt.Println()
	}
}
