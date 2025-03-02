package logic

import printer "github.com/achmad-aditiansah/slice-printer"

func Logic3Pola1(n int) [][]int {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				if i%2 == 0 {
					matrix[i][j] = num
				} else {
					matrix[i][i-j] = num
				}
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola2(n int) [][]int {
	matrix := DynamicSlice(n, n)

	num := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				if i%2 == 0 {
					matrix[i][j] = num
				} else {
					matrix[i][n-1-j+i] = num
				}
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola3(n int) [][]int {
	matrix := DynamicSlice(n, n)

	num := -2

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j <= n-1 {
				if i == 0 || i%2 != 0 {
					num += 3
					if i == 0 {
						matrix[i][j] = num
					} else {
						matrix[i][n-1-i-j] = num
					}
				} else {
					num += 2
					matrix[i][j] = num
				}
			}
		}
	}
	matrix[0][0] = 2
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola4(n int) [][]int {
	matrix := DynamicSlice(n, n)

	num := 1
	for i := 0; i < n; i++ {
		geser := 0
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 1 {
					matrix[i][j] = num
				} else {
					//harusnya 0,8
					//harusnya  yg diisi 6, 7, 8
					matrix[i][n-1-geser] = num
					geser++
				}
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola5(n int) [][]int {
	matrix := DynamicSlice(n, n) // n=9

	num := 1

	for i := 0; i < n; i++ {
		if i <= n/2 { // Untuk baris atas dan tengah
			for j := 0; j <= i; j++ {
				matrix[i][j] = num
				matrix[i][n-1-j] = num
				num += 2
			}
		} else { // Cerminan dari baris sebelumnya
			for j := 0; j < n; j++ {
				matrix[i][j] = matrix[n-1-i][j]
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola6(n int) [][]int {
	matrix := DynamicSlice(n, n)

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if i <= j {
				if i%2 == 0 {
					matrix[i][j] = num
					matrix[n-1-i][j] = num
				} else {
					matrix[i][n-1-j] = num
					matrix[n-1-i][n-1-j] = num
				}
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola7(n int) [][]int {
	matrix := DynamicSlice(n, n)

	nTengah := (n - 1) / 2
	num := 1
	for i := 0; i <= nTengah; i++ {
		for j := 0; j < n; j++ {
			if i+j >= nTengah && j-i <= nTengah {
				if i%2 == 1 {
					matrix[i][j] = num
					matrix[n-1-i][j] = num
				} else {
					matrix[i][n-1-j] = num
					matrix[n-1-i][n-1-j] = num
				}
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola8(n int) [][]int {
	matrix := DynamicSlice(n, n)

	nTengah := (n - 1) / 2
	num := 1
	for i := 0; i <= nTengah; i++ {
		for j := 0; j < n; j++ {
			if i+j >= nTengah && j-i <= nTengah {
				if i%2 == 1 {
					matrix[j][i] = num
					matrix[j][n-1-i] = num
				} else {
					matrix[n-1-j][i] = num
					matrix[n-1-j][n-1-i] = num
				}
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola9(n int) [][]int {
	matrix := DynamicSlice(n, n)

	nTengah := (n - 1) / 2

	for j := 0; j <= nTengah; j++ {
		num := 1
		for i := 0; i <= nTengah; i++ {
			if i+j >= nTengah {
				matrix[j][i] = num
				matrix[n-1-i][j] = num
				matrix[i][n-1-j] = num
				matrix[n-1-i][n-1-j] = num
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola10(n int) [][]int {
	matrix := DynamicSlice(n, n)

	nTengah := (n - 1) / 2

	for j := 0; j <= nTengah; j++ {
		num := 9
		for i := 0; i <= nTengah; i++ {
			if i+j >= nTengah {
				matrix[j][i] = num
				matrix[n-1-i][j] = num
				matrix[i][n-1-j] = num
				matrix[n-1-i][n-1-j] = num
				num -= 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola13(n int) [][]int {
	matrix := DynamicSlice(n, n)

	nTengah := (n - 1) / 2

	for i := 0; i <= nTengah; i++ {
		num := 1
		for j := 0; j <= nTengah; j++ {
			if i+j >= nTengah {
				if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
					matrix[i][j] = num
					matrix[n-1-i][j] = num
					matrix[n-1-i][n-1-j] = num
					matrix[i][n-1-j] = num
				}
			}
			num += 2
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}

func Logic3Pola14(n int) [][]int {
	matrix := DynamicSlice(n, n)

	num := 1
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if j%2 == 0 {
				matrix[i][j] = num
				num += 2
			} else {
				matrix[n-1-i][j] = num
				num += 2
			}
		}
	}
	printer.PrintSlice2D(matrix)
	return matrix
}
