package logic02

func Soal6b(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][n-j-1] = start
				start += 2
			}
		}
	}
	return result
}
