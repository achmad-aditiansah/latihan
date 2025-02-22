package logic

import (
	printer "github.com/achmad-aditiansah/slice-printer"
)

func Logic3Pola1 (n int) {
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
	printer.PrintSlice2D(matrix)
}

func Logic3Pola2 (n int) {
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
	printer.PrintSlice2D(matrix)
}