package logic02

func Create2DArray(n int) [][]int {
	array := make([][]int, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, n)
	}

	return array
}
