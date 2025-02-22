package logic

import (
	printer "github.com/achmad-aditiansah/slice-printer"
)


func DynamicSlice(row, col int) [][]int {

	cell := make([][]int, row)
	for i := range cell {
		cell[i] = make([]int, col)
	}
	return cell
}

func Logic2Pola1 (n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i] = num
		}
		num += 2
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola2 (n int) {
	matrix := DynamicSlice(n, n)

	num := 2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i] = num
		}
		num += 2
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola3 (n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = num
			num += 2
		}
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola4 (n int) {
	matrix := DynamicSlice(n, n)
	
	num := 1 

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = num
			num += 3
		}
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola5 (n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i % 2 == 0 {
				matrix[i][j] = num
			} else {
				matrix[i][n-1-j] = num
			}
			num += 2	
		}
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola6 (n int) {
	matrix := DynamicSlice(n, n)

	num := 1	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i % 2 == 0 {
				matrix[i][j] = num
				num += 3
			} else {
				matrix[i][n-1-j] = num
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola7 (n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				matrix[i][j] = num
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola8(n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i + j == n - 1 {
				matrix[j][i] = num
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola9(n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i == j || i + j == n - 1) {
				matrix[j][i] = num
			}
		}
		num += 2
	}
	printer.PrintSlice2D(matrix)		
}

func Logic2Pola10(n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				matrix[j][i] = num
			}
		}
		num += 2
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola11(n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				matrix[j][i] = num
			}
		}
		num += 2
	}
	printer.PrintSlice2D(matrix)
}

func Logic2Pola12(n int) {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i <= j && (i+j) <= (n-1)) || (i >= j && (i+j) >= (n-1)) {
				matrix[j][i] = num
			}
		}
		num += 2
	}
	printer.PrintSlice2D(matrix)
}