package logic

func Logic3Pola1 (n int) [][]int{
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j ++{
			if i >= j {
				matrix[i][j] = num
			}
			num += 2
		}
	}
	return matrix
}

func Logic3Pola2 (n int) [][]int{
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				matrix[i][j] = num
			}
			num += 2
		}
	}
	return matrix
}